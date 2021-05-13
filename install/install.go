package install

import (
	"fmt"
	d "test/model"
	"test/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Install(c *gin.Context) {

	db := Idb() //先连接数据库，并且初始化数据表
	b := Auser(db)
	if b {
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
		err := db.Create(&userNew).Error
		if err != nil {
			fmt.Println("创建用户失败")
			return false
		} else {
			fmt.Println("成功")
			return true
		}
	} else {
		fmt.Println("后台有用户")
		return false
	}

}

func Idb() *gorm.DB {
	db := d.GetDb()
	db.AutoMigrate(&d.User{})
	db.AutoMigrate(&d.Tp{})
	db.AutoMigrate(&d.View{})
	db.AutoMigrate(&d.Config{})
	return db
}
