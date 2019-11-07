package server

import (
	"github.com/datty258/gorm_project/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		//用户详情
		v1.GET("users/:id", api.ShowUser)
		//所有用户
		v1.GET("users", api.AllUser)
	}
	return r
}
