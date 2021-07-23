package server

import (
	"fmt"
	"goblog/dao"
	"goblog/model"
)

//获取列表,恶心死了gorm- 字段无法进入join关联查询结果...调了我半天.
func GetViewlist(id interface{}, page int) (vi []model.ViewJson) {
	db := dao.MDB.Table("views").Select("views.title,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
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
		fmt.Printf("123%+v", &vi)
	}
	return
}

//获取当前分类下面的10条文章
func Findlist2(id string) (vi []model.ViewJson) {
	fmt.Printf("start")
	db := dao.MDB.Table("views").Select("views.title,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
	db = db.Joins("left join tps on tps.id = views.typeid")
	if id != "0" {
		db = db.Where("typeid = ?", id)
	}
	db.Limit(10).Order("created_at desc").Find(&vi)
	return
}
