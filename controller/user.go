package controller

/*
 * @Description:
 * @Author: longfei
 * @Date: 2021-04-27 09:26:03
 * @LastEditTime: 2021-05-06 15:00:10
 * @LastEditors: longfei
 * @FilePath: \go\controller\user.go
 */

import (
	"fmt"
	d "test/model"
)

type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"` //1:男、2:女
	Pwd    string `json:"pwd"`
}

func TestAdd() {
	user := new(User)
	user.Name = "tome"
	user.Age = 18
	user.Gender = 1
	user.Pwd = "123456"
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
