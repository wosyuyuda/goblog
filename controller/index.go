package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"net/http"

	"fmt"
	d "test/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type view struct {
	gorm.Model
	Type  int
	Title string `gorm:"size:255"`
	Body  string
}

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	u := new(view)
	db.Where("id = ?", id).Find(&u)
	fmt.Println(u)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view":  u,
		"ID":    u.ID,
		"Title": u.Title,
		"Body":  u.Body,
		"Type":  u.Type,
	})
}

//这里是首页
func Views(c *gin.Context) {
	id1 := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	u := new(view)
	li := db.Where("id = ?", id1).Find(&u)
	fmt.Printf("aa%+v\n", li)
	c.HTML(http.StatusOK, "tt.html", gin.H{
		"title": u.Title,
		"id":    u.ID,
		"body":  u.Body,
	})

}

//这里是列表页，后面得加一个分页的才行
func Lists1(c *gin.Context) {
	view2 := new(view)
	views := view2.Findall()
	fmt.Printf("找到的数据是%+v\n", views)
	for _, v := range views {
		fmt.Println(v)
	}
	c.HTML(http.StatusOK, "t1.html", gin.H{
		"list": views,
	})
}

func (view) Findall() (vi []view) {
	db := d.LinkDb() //连接数据库模型
	db.Limit(10).Order("created_at desc").Find(&vi)
	return
}
