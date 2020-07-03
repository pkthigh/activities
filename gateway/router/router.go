package router

import (
	"activities/gateway/handler"

	"github.com/gin-gonic/gin"
)

// NewRouter 新的路由
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())

	// 对外
	v1 := router.Group("/v1")
	{
		activity := v1.Group("/activity")
		{

			activity.GET("/participate", handler.ActivityParticipate)
			activity.POST("/participate", handler.ActivityParticipate)
		}
	}

	// 对内
	backstage := router.Group("/backstage")
	{
		/*
			待上线活动控制
		*/

		// 新增待上线活动
		backstage.POST("/activity", handler.ActivityControl)
		// 更新待上线活动
		backstage.PUT("/activity", handler.ActivityControl)
		// 删除待上线活动
		backstage.DELETE("/activity", handler.ActivityControl)

		/*
			已上线活动控制
		*/

		// 提前下线活动
		backstage.POST("/activity/offline", handler.ActivityOffline)

	}
	return router
}
