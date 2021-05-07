package model

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\model\model.go
 */

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//这里连接数据库，后面可以移到专门的配置文件
func LinkDb() *gorm.DB {
	dsn := "gorm:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func GetDb() (conn *gorm.DB) {
	//这里可以切换成多数据库？
	for {
		conn = dbConn("gorm", "123456", "127.0.0.1", "gorm", 3306)
		if conn != nil {
			break
		}
		fmt.Println("本次未获取到mysql连接")
	}
	return conn
}
