package controller

import (
	"fmt"
	"goblog/dao"
	d "goblog/model"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

//删除分类
func DelType(c *gin.Context) {
	id := c.Param("id")
	db := dao.MDB //连接数据库模型
	err := db.Model(d.Tp{}).Where("id = ?", id).Update("status", "0").Error
	if err != nil {
		fmt.Printf("错误")
	}
	c.JSON(200, gin.H{"msg": "删除成功", "code": 200})
}

func GetTypeNew(id string) (Tp []d.Tp) {
	db := dao.MDB //连接数据库模型
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
	c.JSON(200, gin.H{"msg": "ok", "code": 200, "data": tp})
}

//前端添加跟保存分类
func AddTypes(c *gin.Context) {
	var err error
	data := new(d.Tp)
	err = c.ShouldBind(data)
	//fmt.Printf("这个保存数据传入的数据类型是%T 数据是%+v", data, data)
	//id := c.PostForm("id")
	if data.ID == 0 {
		err = dao.MDB.Create(data).Error
	} else {

		err = dao.MDB.Model(&d.Tp{}).Where("id = ?", data.ID).Updates(&data).Error
	}
	if err != nil {
		fmt.Println("失败")
		c.JSON(400, gin.H{"msg": "失败", "code": 400})
		return
	}
	util.DelAll()
	c.JSON(200, gin.H{"msg": "成功", "code": 200})
}
