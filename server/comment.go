package server

import (
	"goblog/dao"
	"goblog/model"
)

//添加评论
func AddComment(comment *model.Comment) (err error) {
	comment.Status = 2 //设置默认审核
	err = dao.MDB.Create(comment).Error
	return
}

//后台评论列表
func CommentList(page *model.PageList) (err error) {
	err = dao.Page(page).Find(&page.Comment).Count(&page.Sum).Error
	return
}

//文章所属评论
func ViewComment(view *model.View) (err error) {
	err = dao.MDB.Where("vid = ? AND status = ?", view.ID, 1).Find(&view.Comment).Error
	return
}

//评论的审核与删除
func F审核(评论 *model.Comment) (错误 error) {
	错误 = dao.MDB.Where("id = ?", 评论.ID).Update("status", 评论.Status).Error
	return
}
