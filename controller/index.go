package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"net/http"

	d "test/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type view struct {
	gorm.Model
	Typeid int    //分类的ID，关联
	Title  string `gorm:"size:255"` //标题
	Body   string //详细的内容
	Click  int    `gorm:"size:10"`  //点击量
	Pic    string `gorm:"size:255"` //文章的缩略图
	Tps    Tp     `json:"tps" gorm:"FOREIGNKEY:Typeid;"`
}

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	u := new(view)
	db.Where("id = ?", id).Find(&u)
	db.Model(&u).Preload("Tps").Find(&u)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view": u,
	})
}

//这里是列表页
func Views(c *gin.Context) {
	id1 := c.Param("id")
	v1 := new(view)
	list := v1.Findlist(id1)
	c.HTML(http.StatusOK, "t1.html", gin.H{
		"title": "这个是后台设置的标题呀",
		"list":  list,
	})

}

//这里查询列表
func (view) Findlist(id string) (vi []view) {
	db := d.LinkDb() //连接数据库模型

	if id == "0" {
		db.Limit(10).Order("created_at desc").Preload("Tps").Find(&vi)
		//下面是原来的
		//db.Limit(10).Order("created_at desc").Find(&vi)
		//db.Model(&vi).Preload("Tps").Find(&vi)
	} else {
		db.Where("typeid = ?", id).Limit(10).Order("created_at desc").Preload("Tps").Find(&vi)
	}
	return
}

//这里是首页
func Lists1(c *gin.Context) {
	view2 := new(view)
	views := view2.Findlist("0")
	c.HTML(http.StatusOK, "t1.html", gin.H{
		"list": views,
	})
}
