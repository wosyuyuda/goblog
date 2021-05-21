package util

import (
	d "goblog/model"
)

//循环的给文章添加默认图片
func Imgsrc(v []d.View) []d.View {
	for i := 0; i < len(v); i++ {
		if len(v[i].Pic) == 0 {
			v[i].Pic = "/static/no-images.jpg"
		}
	}
	return v
}

//获取分类下有多少文章，或者总共有多少
func GetTypeCount(id string) int64 {
	var i int64
	db := d.LinkDb()
	if id == "0" {
		db.Find(&d.View{}).Count(&i)
	} else {
		db.Model(&d.View{}).Where("typeid = ?", id).Count(&i)
	}
	return i
}

//获取当前文章信息，如果up为1那么点击数+1
func GetView(id string, up int) *d.View {
	db := d.LinkDb()
	u := new(d.View)
	db.Where("id = ?", id).Find(&u)
	db.Model(&u).Preload("Tps").Find(&u)
	if up == 1 {
		db.Model(&u).Update("click", u.Click+1) //点击量加1
	}
	return u
}
