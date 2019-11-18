package api

import (
	"github.com/datty258/gorm_project/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 随机返回一条毒鸡汤
func ShowSoul(c *gin.Context) {
	var soul model.Soul
	res := model.DB.Order(gorm.Expr("rand()")).Find(&soul)
	if res == nil {
		panic("找不到结果")
	} else {
		c.JSON(200, soul.Content)
	}

}
