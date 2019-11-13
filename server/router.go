package server

import (
	"os"

	"github.com/datty258/gorm_project/api"
	"github.com/datty258/gorm_project/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))

	v1 := r.Group("/api/v1")
	{
		//用户详情
		v1.GET("users/:id", api.ShowUser)
		//所有用户
		v1.GET("users", api.AllUser)
	}
	return r
}
