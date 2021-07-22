package controller

import (
	"fmt"
	"goblog/dao"
	d "goblog/model"

	"github.com/gin-gonic/gin"
)

//这个用来处理后台设置的问题,这些设置平常用的比较多,如果有redis那么存进缓存会好一些
func GetConfig(c *gin.Context) {
	db := dao.MDB
	n := new(d.Config)
	db.Where("name = ?", c.Param("name")).Find(&n)
	//后面加一个统一的返回json接口
	c.JSON(200, gin.H{"msg": "数据获取成功", "code": 200, "data": n})
}

//新增与保存都用这个接口
func SetConfig(c *gin.Context) {
	var data []d.Config
	c.ShouldBind(&data)
	fmt.Printf("传过来的数据是%+v", data) //接收后看数据是否正常,再判断是新增还是保存
}
