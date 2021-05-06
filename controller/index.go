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
)

type view struct {
	Id    int
	Type  int
	Title string
	Body  string
}

//这里加一个接收前端数据的再返回数据就好啦
func AddView(c *gin.Context) {
	view1 := new(view)
	view1.Type = 1
	view1.Title = "这个是标题的测试"

	conn := d.GetDb()
	conn.AutoMigrate(&view{})
	err := conn.Create(view1).Error
	if err != nil {
		fmt.Println("创建失败")
	}
	fmt.Println("创建成功")
}

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	u := new(view)
	db.Where("id = ?", id).Find(&u)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view": u,
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
		"id":    u.Id,
		"body":  u.Body,
	})

}

//这里是列表页
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
	db.Find(&vi)
	return
}
