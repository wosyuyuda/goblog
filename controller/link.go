package controller

import (
	"goblog/dao"
	"goblog/model"
	"goblog/server"

	"github.com/gin-gonic/gin"
)

//添加编辑友情链接
func AddLink(c *gin.Context) {
	var link model.Link
	err := c.ShouldBind(&link)
	if err != nil {
		server.Fail(c)
		return
	}
	if link.ID == 0 {
		err = dao.MDB.Create(&link).Error
	} else {
		err = dao.MDB.Where("id = ?", link.ID).Updates(&link).Error
	}
	server.IfRes(err, c)
}

//获取友链列表
func GetLink(c *gin.Context) {
	var link []model.Link
	err := dao.MDB.Find(&link).Error
	server.ResDataError(link, err, c)
}
