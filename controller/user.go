package controller

/*
 * @Description:
 * @Author: longfei
 * @LastEditors: Please set LastEditors
 * @FilePath: \go\controller\user.go
 */

import (
	"fmt"
	d "test/model"
	"test/util"

	"github.com/gin-gonic/gin"
)

type User struct {
	d.User
}

func AddU(c *gin.Context) {
	user := new(User)
	user.Name = "longfei"
	user.Age = 18
	user.Gender = 1
	pwd := "123456"
	user.Pwd = util.Md5(pwd)
	user.AddUser() //user.id = 1，添加之后user中的id会变成数据库中生成的值
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
