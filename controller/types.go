package controller

import (
	"fmt"
	d "test/model"

	"github.com/gin-gonic/gin"
)

func GetTypeNew(id string) (Tp []d.Tp) {
	db := d.LinkDb() //连接数据库模型
	switch id {
	case "0":
		db.Find(&Tp)
	case "-1":
		db.Find(&Tp)
	default:
		db.Where("id = ?", id).Find(&Tp)
	}
	return
}

//前端获取分类
func Gt(c *gin.Context) {
	tp := GetTypeNew("0")
	c.JSON(200, gin.H{"msg": "验证码错误", "code": 200, "data": tp})
}

func AddTypes(c *gin.Context) {
	tp1 := new(d.Tp)
	tp1.Name = c.PostForm("name")
	conn := d.GetDb()
	conn.AutoMigrate(&d.Tp{})
	err := conn.Create(tp1).Error
	if err != nil {
		fmt.Println("创建失败")
		c.JSON(400, gin.H{"msg": "创建失败", "code": 400})
	}
	c.JSON(200, gin.H{"msg": "创建成功", "code": 200})
}

func SaveType(c *gin.Context) {

}
