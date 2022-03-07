package controller

import (
	"goblog/model"
	"goblog/server"

	"github.com/gin-gonic/gin"
)

//添加评论的接口,害又得加后台了..这个先放着吧.
func AddComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBind(&comment)
	if err != nil {
		server.RData(err, c)
		return
	}
	err = server.AddComment(&comment)
	server.RData(err, c, &comment)
}

//评论列表
func ListComment(c *gin.Context) {
	var page model.PageList
	err := c.ShouldBind(&page)
	if err != nil {
		server.RData(err, c)
		return
	}
	err = server.CommentList(&page)
	server.RData(err, c, &page)
}

//评论的审核与删除
func F审核(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBind(&comment)
	if err != nil {
		server.RData(err, c)
		return
	}
	err = server.F审核(&comment)
	server.RData(err, c)
}
