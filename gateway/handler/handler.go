package handler

import (
	"activities/common/errs"
	"activities/gateway"
	"activities/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ActivityParticipate 活动参与
func ActivityParticipate(c *gin.Context) {
	type ParamModel struct {
		ActivityID int64 `form:"activity_id" json:"activity_id"`
	}
	if c.Request.Method == "GET" {
		param := &ParamModel{}
		if err := c.ShouldBindQuery(&param); err != nil {
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

		if result := service.Service.IsMeetParticipationConditions(param.ActivityID, 0); result {
			c.JSON(200, gateway.Successful(nil))
			return
		} else if !result {
			c.JSON(200, &gateway.ResultInfo{
				Code: errs.NotMeetParticipationConditions.Int(),
				Msg:  "不满足该活动参与条件",
			})
			return
		}
	}

	if c.Request.Method == "POST" {
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
		if err := service.Service.Participation(param.ActivityID, 0); err != nil {
			c.JSON(200, gateway.ResultInfo{
				Code: errs.ParticipationFailure.Int(),
				Msg:  "参与活动失败",
			})
			return
		}
	}
	return
}

// ActivityControl 活动控制
func ActivityControl(c *gin.Context) {
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

	// 新增待上线活动
	if c.Request.Method == "POST" {
		if err := service.Service.AddActivity(param.ActivityID); err != nil {
			return
		}
	}

	// 更新待上线活动
	if c.Request.Method == "PUT" {
		if err := service.Service.UpdateActivity(param.ActivityID); err != nil {
			return
		}
	}

	// 删除待上线活动
	if c.Request.Method == "DELETE" {
		if err := service.Service.DelActivity(param.ActivityID); err != nil {
			return
		}
	}

	c.JSON(200, gateway.Successful(nil))
	return
}

// ActivityOffline 活动下线
func ActivityOffline(c *gin.Context) {
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
}
