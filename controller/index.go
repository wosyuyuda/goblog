package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"net/http"
	"strconv"

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

//这里加一个接收前端数据的再返回数据就好啦，应该再加一个是否登陆判断
func AddView(c *gin.Context) {
	view1 := new(view)
	type1 := c.PostForm("type") //这里分类还要转成int类型，真麻烦,好像直接用string还方便些
	viewType, _ := strconv.Atoi(type1)
	view1.Type = viewType

	view1.Title = c.PostForm("title")
	view1.Body = c.PostForm("body")
	fmt.Printf("传过来的标题是：%s 密码是：%s", c.PostForm("body"), c.PostForm("title"))
	//fmt.Println(view1)
	conn := d.GetDb()
	conn.AutoMigrate(&view{})
	err := conn.Create(view1).Error
	if err != nil {
		fmt.Println("创建失败")
	}
	fmt.Println("创建成功")
	c.JSON(200, gin.H{"msg": "创建成功", "code": 200})
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
		"id":    u.ID,
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
