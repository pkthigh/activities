package handler

import (
	"activities/common"
	"activities/common/errs"
	"activities/gateway"
	"activities/library/logger"
	"activities/library/storage"
	"activities/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Activity 活动
func Activity(c *gin.Context) {
	type ParamModel struct {
		ActivityID int64 `json:"activity_id"`
	}
	param := &ParamModel{}
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		c.JSON(200, gateway.ResultInfo{
			Code: errs.IllegalParameter.Int(),
			Msg:  err.Error(),
		})
		return
	}
	if param.ActivityID == 0 {
		c.JSON(200, gateway.ResultInfo{
			Code: errs.MissingParameter.Int(),
			Msg:  "参数缺失",
		})
		return
	}

	db, ok := storage.GetSQLDB(common.ActivityDsn)
	if !ok {
		c.JSON(200, gateway.ResultInfo{
			Code: errs.DatabaseError.Int(),
			Msg:  "获取ActivityDsn数据库游标错误",
		})
		return
	}
	// 新增活动待上线
	if c.Request.Method == "POST" {

		activity := &models.AcActivity{}
		if err := db.Where("`id` = ?", param.ActivityID).First(activity).Error; err != nil {
			logger.ErrorF("Activity POST First AcActivity %v error: %v", param.ActivityID, err)
			c.JSON(200, gateway.ResultInfo{
				Code: errs.DatabaseError.Int(),
				Msg:  err.Error(),
			})
			return
		}

		if activity.Status != common.Not.Int() {
			logger.ErrorF("Activity POST ActivityID: %v Status: %v", param.ActivityID, activity.Status)
			c.JSON(200, gateway.ResultInfo{
				Code: errs.IllegalParameter.Int(),
				Msg:  "非法活动状态",
			})
			return
		}

		c.JSON(200, gateway.ResultInfo{
			Code: errs.Default.Int(),
		})
		return
	}

	// 更新待上线活动
	if c.Request.Method == "PUT" {

	}
	// 下线活动
	if c.Request.Method == "DELETE" {

	}
}
