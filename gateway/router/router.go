package router

import "github.com/gin-gonic/gin"

// NewRouter 新的路由
func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}
