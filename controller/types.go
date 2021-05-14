package controller

import (
	"fmt"
	"strconv"
	d "test/model"

	"github.com/gin-gonic/gin"
)

//删除分类
func DelType(c *gin.Context) {
	id := c.Param("id")
	db := d.LinkDb() //连接数据库模型
	err := db.Model(d.Tp{}).Where("id = ?", id).Update("status", "0").Error
	if err != nil {
		fmt.Printf("错误")
	}
	fmt.Println(id)
}

func GetTypeNew(id string) (Tp []d.Tp) {
	db := d.LinkDb() //连接数据库模型
	switch id {
	case "0":
		db.Where("status = ?", "1").Find(&Tp)
	case "-1":
		db.Where("status = ?", "1").Find(&Tp)
	default:
		db.Where("status = ?", "1").Where("id = ?", id).Find(&Tp)
	}
	return
}

//前端获取分类
func Gt(c *gin.Context) {
	tp := GetTypeNew("0")
	c.JSON(200, gin.H{"msg": "验证码错误", "code": 200, "data": tp})
}

//前端添加跟保存分类
func AddTypes(c *gin.Context) {
	var err error
	data := new(d.Tp)
	data.Name = c.PostForm("name")
	data.Info = c.PostForm("info")
	leve, _ := strconv.Atoi(c.PostForm("level"))
	data.Level = leve
	fmt.Printf("这个保存数据传入的数据类型是%T 数据是%+v", data, data)
	id := c.PostForm("id")
	conn := d.GetDb()
	conn.AutoMigrate(&d.Tp{})
	if id != "0" {
		err = conn.Model(&d.Tp{}).Where("id = ?", id).Updates(&data).Error
	} else {
		err = conn.Create(data).Error
	}

	if err != nil {
		fmt.Println("失败")
		c.JSON(400, gin.H{"msg": "失败", "code": 400})
	}
	c.JSON(200, gin.H{"msg": "成功", "code": 200})
}
