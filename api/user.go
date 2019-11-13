package api

import (
	"github.com/datty258/gorm_project/model"
	"github.com/datty258/gorm_project/serializer"
	"github.com/gin-gonic/gin"
)

// ShowUser 查看固定id用户
// 包括用户包含的对应关系的信息
func ShowUser(c *gin.Context) {
	var user model.User
	res := model.DB.First(&user, c.Param("id")).
		Related(&user.Languages, "languages").
		Related(&user.Addresses).
		Related(&user.Tel)

	re := serializer.BuildUserResponse(user)

	if res == nil {
		panic("找不到该用户")
	} else {
		c.JSON(200, re)
	}

}

// AllUser 所有用户列表
func AllUser(c *gin.Context) {
	var users = []model.User{}
	res := model.DB.Order("id desc").Find(&users)
	if res == nil {
		panic("找不到用户")
	} else {
		c.JSON(200, res)
	}
}
