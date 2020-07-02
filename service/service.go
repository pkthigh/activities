package service

import (
	"activities/common"
	"activities/common/errs"
	"activities/library/logger"
	"activities/library/storage"
	"activities/models"
	"activities/service/framework"
	"sync"
)

// Service 活动服务
var Service *ActivitiesService

func init() {
	Service = &ActivitiesService{msgs: make(chan string, 1024)}
	go Service.listening()
}

// ActivitiesService 活动服务
type ActivitiesService struct {
	msgs chan string
	// key: common.ActivityType val: framework.Activity
	activities sync.Map
}

// listening 消费Nats的消息
func (srv *ActivitiesService) listening() {
	for {
		userid := <-srv.msgs
		srv.activities.Range(func(k, v interface{}) bool {
			activity := v.(framework.Activity)
			if result := activity.Judge(userid); result {
				// TODO: 通知用户
			}
			return true
		})
	}
}

// PostMsg 投递消息
func (srv *ActivitiesService) PostMsg(userid string) {
	srv.msgs <- userid
}

// AddActivity 添加活动
func (srv *ActivitiesService) AddActivity(id int64) error {
	activity, err := srv.getActivity(id)
	if err != nil {
		logger.ErrorF("AddActivity getActivity error: %v", err)
		return errs.DatabaseError.Error()
	}
	details, err := srv.getDetails(id)
	if err != nil {
		logger.ErrorF("AddActivity getActivity error: %v", err)
		return errs.DatabaseError.Error()
	}
	var ty = common.ActivityType(activity.AcType)

	if _, ok := srv.activities.Load(ty); ok {
		logger.ErrorF("AddActivity this activity type is already planned")
		return errs.DatabaseError.Error()
	}
	switch ty {
	case common.HandNum:
		srv.activities.Store(common.HandNum, framework.NewHandNumActivity(activity, details))

	default:
		logger.ErrorF("AddActivity unknown activity type: %v", ty)
		return errs.UnknownActivityType.Error()
	}
	return nil
}

// UpdateActivity 更新活动
func (srv *ActivitiesService) UpdateActivity(id int64) error {
	activity, err := srv.getActivity(id)
	if err != nil {
		logger.ErrorF("UpdateActivity getActivity error: %v", err)
		return errs.DatabaseError.Error()
	}
	details, err := srv.getDetails(id)
	if err != nil {
		logger.ErrorF("UpdateActivity getActivity error: %v", err)
		return errs.DatabaseError.Error()
	}

	var ty = common.ActivityType(activity.AcType)
	if val, ok := srv.activities.Load(ty); !ok { // 当前没有该类型活动在排期
		logger.ErrorF("UpdateActivity this activity type is not currently set: %v", ty)
		return errs.ThisActivityTypeNotCurrentlySet.Error()

	} else if ok {
		var ac = val.(framework.Activity)
		if ac.ID() == id {
			switch ty {
			case common.HandNum:
				srv.activities.Store(ty, framework.NewHandNumActivity(activity, details))
			}
		} else {
			logger.ErrorF("UpdateActivity activity id not match, now: %v param: %v", ac.ID(), id)
			return errs.ActivityIDNotMatch.Error()
		}
	}

	return nil
}

// DelActivity 删除活动
func (srv *ActivitiesService) DelActivity(id int64) error {
	var ty common.ActivityType = -1
	srv.activities.Range(func(k, v interface{}) bool {
		if v.(framework.Activity).ID() == id {
			ty = k.(common.ActivityType)
			return false
		}
		return true
	})

	if ty == -1 { // 当前没有该类型活动在排期
		logger.ErrorF("DelActivity this activity type is not currently set: %v", ty)
		return errs.ThisActivityTypeNotCurrentlySet.Error()

	} else if ty > 0 {
		srv.activities.Delete(ty)
	}
	return nil
}

// OffActivity 下线活动
func (srv *ActivitiesService) OffActivity(id int64) error {
	var ty common.ActivityType = -1
	srv.activities.Range(func(k, v interface{}) bool {
		if v.(framework.Activity).ID() == id {
			ty = k.(common.ActivityType)
			return false
		}
		return true
	})

	if ty == -1 { // 当前没有该类型活动在排期
		logger.ErrorF("OffActivity this activity type is not currently set: %v", ty)
		return errs.ThisActivityTypeNotCurrentlySet.Error()
	}

	val, _ := srv.activities.Load(ty)
	var ac = val.(framework.Activity)
	if ac.Status() != common.Ongoing {
		logger.ErrorF("OffActivity this activity status not ongoing, now: %v", ac.Status())
		return errs.ActivityStatusNotMatch.Error()
	}
	db, _ := storage.GetSQLDB(common.ActivityDsn) // 更新数据库
	if err := db.Where("`id = ?`", id).Update("status", common.Offline).Error; err != nil {
		logger.ErrorF("OffActivity Update status error: %v", err)
		return errs.DatabaseError.Error()
	}
	srv.activities.Delete(ty)
	return nil
}

// getActivity 从数据库获取活动信息
func (srv *ActivitiesService) getActivity(id int64) (*models.AcActivity, error) {
	db, _ := storage.GetSQLDB(common.ActivityDsn)

	var activity *models.AcActivity
	if err := db.Where("`id` = ?", id).First(activity).Error; err != nil {
		return nil, err
	}

	return activity, nil
}

// getDetails 从数据库获取活动细节
func (srv *ActivitiesService) getDetails(id int64) ([]*models.AcActivityDetail, error) {
	db, _ := storage.GetSQLDB(common.ActivityDsn)

	var details []*models.AcActivityDetail
	if err := db.Where("`activity_id` = ?", id).Find(&details).Error; err != nil {
		return nil, err
	}
	return details, nil
}
