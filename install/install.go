package install

import (
	"fmt"
	d "test/model"
	"test/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Install(c *gin.Context) {
	db := d.GetDb()
	Idb(db)        //初始化数据库表
	b := Auser(db) //添加第一个用户
	e := Atype(db) //添加第一个分类
	f := Aview(db) //添加第一篇文章
	if b && e && f {
		c.JSON(200, gin.H{"msg": "初始化成功", "code": 200})
	} else {
		c.JSON(200, gin.H{"msg": "初始化失败", "code": 400})
	}

}

//添加用户
func Auser(db *gorm.DB) bool {
	var user d.User
	db.First(&user)
	if user.Id == 0 {
		userNew := new(d.User)
		userNew.Name = "longfei"
		userNew.Pwd = util.Md5("123456") //md5加密一下
		return db.Create(&userNew).Error == nil
	} else {
		fmt.Println("后台有用户")
		return false
	}
}

//添加一个分类
func Atype(db *gorm.DB) bool {
	tp := new(d.Tp)
	tp.Name = "最新"
	tp.Info = "第一个分类"
	return db.Create(&tp).Error == nil
}

func Aview(db *gorm.DB) bool {
	view := new(d.View)
	view.Title = "这是第goblog的第一篇文章"
	view.Typeid = 1
	view.Content = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg"
	view.Body = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg"
	return db.Create(&view).Error == nil
}

//初始化数据库表文件
func Idb(db *gorm.DB) {
	db.AutoMigrate(&d.User{})
	db.AutoMigrate(&d.Tp{})
	db.AutoMigrate(&d.View{})
	db.AutoMigrate(&d.Config{})
}
