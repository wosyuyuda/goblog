package model

/*
 * @Description:
 * @Author: longfei
 * @Date: 2021-04-25 09:18:47
 * @LastEditTime: 2021-04-25 17:46:25
 * @LastEditors: longfei
 * @FilePath: \go\model\model.go
 */

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//这里连接数据库
func LinkDb() *gorm.DB {

	dsn := "gorm:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//文章详情数据
func Views() {
	type View struct {
		gorm.Model
		Title string
		Body  string
	}
	db := LinkDb()
	fmt.Println("123")
	db.AutoMigrate(&View{})
}
