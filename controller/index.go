package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"net/http"
	"strconv"

	d "test/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type view struct {
	gorm.Model
	Typeid  int    `gorm:"size:10"`  //分类的ID，关联
	Title   string `gorm:"size:255"` //标题
	Body    string //详细的内容
	Click   int    `gorm:"size:10"`                       //点击量
	Tuijian uint   `gorm:"size:1"`                        //是否为推荐
	Swiper  uint   `gorm:"size:1"`                        //是否为轮播图
	Pic     string `gorm:"size:255"`                      //文章的缩略图
	Tps     Tp     `json:"tps" gorm:"FOREIGNKEY:Typeid;"` //这里放分类信息types
}

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	u := new(view)
	newList := u.Findlist("0")
	tuijian := u.Findlist("-1")
	db.Where("id = ?", id).Find(&u)
	db.Model(&u).Preload("Tps").Find(&u)
	types1 := new(Tp)
	tp := types1.GetType("0")
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view":    u,
		"types":   tp,
		"newlist": newList,
		"tuijian": tuijian,
	})
}

//这里是列表页
func Views(c *gin.Context) {
	id1 := c.Param("id")
	v1 := new(view)
	list := v1.Findlist(id1)
	newList := v1.Findlist("0")
	tuijian := v1.Findlist("-1")
	types1 := new(Tp)
	tp := types1.GetType("0")
	c.HTML(http.StatusOK, "list.html", gin.H{
		"list":     list,
		"typeinfo": list[0].Tps,
		"types":    tp,
		"newlist":  newList,
		"tuijian":  tuijian,
	})
}

//这里是首页
func Lists1(c *gin.Context) {
	view2 := new(view)
	views := view2.Findlist("0")
	tuijian := view2.Findlist("-1")
	types1 := new(Tp)
	tp := types1.GetType("0")
	//获取首页每个栏目的列表
	tt := []Tp{}
	for _, v := range tp {
		v.Views = view2.Findlist1(strconv.Itoa(int(v.ID)))
		tt = append(tt, v)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"list":    views,
		"types":   tp,
		"tuijian": tuijian,
		"tt":      tt,
	})
}

//这里查询列表,0获取全部的，-1是推荐，-2是轮播
func (view) Findlist(id string) (vi []view) {
	db := d.LinkDb() //连接数据库模型
	switch id {
	case "0":
		db.Limit(10).Order("created_at desc").Preload("Tps").Find(&vi)
	case "-1":
		db.Where("tuijian = ?", 1).Limit(10).Order("created_at desc").Preload("Tps").Find(&vi)
	default:
		db.Where("typeid = ?", id).Limit(10).Order("created_at desc").Preload("Tps").Find(&vi)
	}
	return
}
func (view) Findlist1(id string) (vi []view) {
	db := d.LinkDb() //连接数据库模型
	db.Where("typeid = ?", id).Limit(10).Order("created_at desc").Find(&vi)
	return
}
