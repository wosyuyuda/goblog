package controller

/*
 * @Description:
 * @Author: longfei
 * @LastEditors: Please set LastEditors
 * @FilePath: \go\controller\user.go
 */

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	d "test/model"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `gorm:"size:255" json:"name"`
	Age    int    `gorm:"size:2" json:"age"`
	Gender int    `gorm:"size:2" json:"gender"` //1:男、2:女
	Pwd    string `gorm:"size:255" json:"pwd"`
}

func AddU(c *gin.Context) {
	user := new(User)
	user.Name = "tome"
	user.Age = 18
	user.Gender = 1
	pwd := "123456"

	h := md5.New()
	h.Write([]byte(pwd))
	user.Pwd = hex.EncodeToString(h.Sum(nil)) //md5加密一道
	user.AddUser()                            //user.id = 1，添加之后user中的id会变成数据库中生成的值
}

func (user *User) AddUser() {
	conn := d.GetDb()
	fmt.Println("进入创建")
	conn.AutoMigrate(&User{})
	err := conn.Create(user).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}

//修改数据
/* func (user *User) UpdateUser() {
	conn := d.GetDb()

	err := conn.Model(&User{}).Update(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
} */

//删除数据
func (user *User) DelUser() {
	conn := d.GetDb()
	err := conn.Delete(user).Error
	if err != nil {
		fmt.Println("删除失败")
	}
}
