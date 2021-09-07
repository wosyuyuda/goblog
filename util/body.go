package util

import (
	"fmt"
	"goblog/dao"
	d "goblog/model"
	"time"

	"gorm.io/gorm"
)

//循环的给文章添加默认图片
func Imgsrc(v []d.ViewJson) []d.ViewJson {
	for i := 0; i < len(v); i++ {
		if len(v[i].Pic) == 0 {
			v[i].Pic = "/static/no-images.jpg"
		}
		tm := time.Unix(int64(v[i].CreatedAt), 0)
		v[i].Ctime = tm.Format("2006-01-02 15:04:05")
	}
	return v
}

//获取分类下有多少文章，或者总共有多少
func GetTypeCount(id string) int64 {
	var i int64
	db := dao.MDB
	if id == "0" {
		db.Find(&d.View{}).Count(&i)
	} else {
		db.Model(&d.View{}).Where("typeid = ?", id).Count(&i)
	}
	return i
}

//获取当前文章信息，如果up为1那么点击数+1
func GetView(id string, up int) (dd d.View) {
	err := GetViewCache(id, &dd)
	if err != nil {
		if id != "0" {
			dao.MDB.Model(&d.View{}).Where("id = ?", id).Find(&dd)
		} else {
			dao.MDB.Model(&d.View{}).Where("status = ?", 3).Find(&dd)
		}
		fmt.Println(dd)
		if dd.ID != 0 {
			SetViewCache(&dd)
		}
	}
	if up == 1 {
		dao.MDB.Model(&d.View{}).Where("id = ?", id).UpdateColumn("click", gorm.Expr("click + ?", 1)) //点击量加1
	}
	return
}
