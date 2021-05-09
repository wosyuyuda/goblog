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
	v1 := new(view)
	list := v1.Findlist(id1)
	c.HTML(http.StatusOK, "t1.html", gin.H{
		"list": list,
	})

}

func (view) Findlist(id string) (vi []view) {
	db := d.LinkDb() //连接数据库模型
	if id == "0" {
		db.Limit(10).Order("created_at desc").Find(&vi)
	} else {
		db.Where("type = ?", id).Limit(10).Order("created_at desc").Find(&vi)
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
