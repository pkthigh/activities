package framework

import (
	"activities/common"
	"activities/library/storage"
	"activities/models"
	"strconv"
)

// HandNumActivity 手数活动
type HandNumActivity struct {
	Activity *models.AcActivity
	Details  []*models.AcActivityDetail
}

// NewHandNumActivity 实现抽象活动类的手数活动
func NewHandNumActivity(activity *models.AcActivity, details []*models.AcActivityDetail) Activity {
	return &HandNumActivity{
		Activity: activity,
		Details:  details,
	}
}

// ID 活动ID
func (activity *HandNumActivity) ID() int64 {
	return activity.Activity.ID
}

// Type 活动类型
func (activity *HandNumActivity) Type() common.ActivityType {
	return common.ActivityType(activity.Activity.AcType)
}

// Status 活动状态
func (activity *HandNumActivity) Status() common.ActivityStatus {
	return common.ActivityStatus(activity.Activity.Status)
}

// Judge 判断玩家是否满足参与活动的条件
func (activity *HandNumActivity) Judge(userid string) bool {
	// 已抽奖
	store, _ := storage.GetRdsDB(common.HandNumActivity)
	if store.Exists(userid).Val() {
		return false
	}
	record, _ := storage.GetRdsDB(common.HandOverRecordStore)

	count, _ := strconv.Atoi(record.Get(userid).Val())
	if count >= activity.Activity.HandNum {
		return true
	}

	return false
}

// Participate 参与活动
func (activity *HandNumActivity) Participate(userid string) {
	if !activity.Judge(userid) {
		return
	}
	// TODO: 进行活动
}
