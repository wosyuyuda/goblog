package dao

import (
	"fmt"
	"goblog/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MDB *gorm.DB

//这个自动 连接数据库
func init() {
	//初始化数据库mysql
	var err error
	con := config.Configv

	connArgs := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		con.GetString("db.user"),
		con.GetString("db.pwd"),
		con.GetString("db.host"),
		con.GetInt("db.port"),
		con.GetString("db.db"))

	//fmt.Println("读取配置文件: ", connArgs)
	MDB, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		log.Fatal("connect MDB failed:", err.Error())
		return
	}
	sqlDB, _ := MDB.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100) //最大连接数
	sqlDB.SetMaxIdleConns(10)  //最大空闲连接数
	// SetConnMaxLifetime 设置了连接可复用的最大时间。现在是一小时? time.Minute*5
	sqlDB.SetConnMaxLifetime(time.Second * 30)

}
