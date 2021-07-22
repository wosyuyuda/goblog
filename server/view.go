package server

import (
	"fmt"
	"goblog/dao"
	"goblog/model"
)

//获取列表
func GetViewlist(id interface{}, page int) (vi []model.View) {
	db := dao.MDB.Table("views").Select("views.*, tps.Name as Name")
	db = db.Joins("left join tps on tps.id = views.typeid")
	num := 10 //一页默认10条
	if page < 1 {
		page = 1
	}
	page -= 1 //过滤掉的页面要减 1
	order := "created_at desc"
	switch id {
	case "0":
		db.Limit(num).Offset(page * num).Order(order).Find(&vi)
	case "-1":
		db.Where("tuijian = ?", 1).Limit(num).Order(order).Find(&vi)
	case "-2":
		db.Where("swiper = ?", 1).Limit(num / 2).Order(order).Find(&vi)
	case "-3":
		db.Limit(num).Offset(page * num).Order("click desc").Find(&vi)
	case "-4":
		db.Where("tuijian = ?", 1).Limit(3).Order(order).Find(&vi)
	default:
		db.Where("typeid = ?", id).Limit(num).Offset(page * num).Order(order).Find(&vi)
	}
	return
}

//获取当前分类下面的10条文章
func Findlist2(id string) (vi []model.View) {
	fmt.Printf("start")
	db := dao.MDB.Table("views").Select("views.*, tps.Name as Name")
	db = db.Joins("left join tps on tps.id = views.typeid")
	if id != "0" {
		db = db.Where("typeid = ?", id)
	}
	db.Limit(10).Order("created_at desc").Scan(&vi)
	return
}
