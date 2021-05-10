package controller

import (
	"fmt"
	d "test/model"

	"github.com/gin-gonic/gin"
)

type Tp struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"size:255"`
	views []view `gorm:"foreignKey:Typeid;references:ID"`
}

//-1是获取首页带文章的
func (Tp) GetType(id string) (Tp []Tp) {
	db := d.LinkDb() //连接数据库模型
	switch id {
	case "0":
		db.Find(&Tp)
	case "-1":
		db.Debug().Preload("view").Find(&Tp)
	default:
		db.Where("id = ?", id).Find(&Tp)
	}
	return
}

func AddTypes(c *gin.Context) {
	tp1 := new(Tp)
	tp1.Name = c.PostForm("name")
	conn := d.GetDb()
	conn.AutoMigrate(&Tp{})
	err := conn.Create(tp1).Error
	if err != nil {
		fmt.Println("创建失败")
		c.JSON(400, gin.H{"msg": "创建失败", "code": 400})
	}
	c.JSON(200, gin.H{"msg": "创建成功", "code": 200})
}
