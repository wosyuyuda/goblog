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

//下面直接转过来不行
type View struct {
	gorm.Model
	Typeid  int    `gorm:"size:10"  json:"typeid" form:"typeid"`   //分类的ID，关联
	Title   string `gorm:"size:255" json:"title" form:"title"`     //标题
	Body    string `json:"body" form:"body"`                       //详细的内容
	Click   int    `gorm:"size:10" json:"click"`                   //点击量
	Tuijian uint   `gorm:"size:1" json:"tuijian" form:"tuijian"`   //是否为推荐
	Swiper  uint   `gorm:"size:1" json:"swiper" form:"swiper"`     //是否为轮播图
	Pic     string `gorm:"size:255" json:"pic" form:"pic"`         //文章的缩略图
	Content string `gorm:"size:500" json:"content" form:"content"` //文章的简介
	Status  uint   `gorm:"size:1" json:"status"`                   //文章状态，0删除，1正常
	Tps     Tp     `json:"tps" gorm:"FOREIGNKEY:Typeid;"`          //这里放分类信息types
}

//分类表
type Tp struct {
	ID     uint   `gorm:"primarykey" json:"id" form:"id"`
	Name   string `gorm:"size:255;" form:"name" json:"name"`               //名称
	Level  int    `gorm:"size:10;default:0;" form:"level" json:"level"`    //分类排序
	Info   string `gorm:"size:255;" form:"info" json:"info"`               //介绍
	Status string `gorm:"size:1;default:'1';" form:"status" json:"status"` //状态,默认1,0表示删除
	Views  []View `gorm:"-"`
}

//系统设置
type Config struct {
	ID      uint   `gorm:"primarykey"  json:"id" form:"id"`        //用户ID
	Name    string `gorm:"size:255"  json:"name" form:"name"`      //设置的名称
	Type    int    `gorm:"size:10"  json:"type" form:"type"`       //设置所属分类
	Content string `gorm:"size:500" json:"content" form:"content"` //设置的介绍
	Value   string `gorm:"size:500" json:"value" form:"value"`     //系统设置值
}

//用户表
type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `gorm:"size:255" json:"name"`
	Age    int    `gorm:"size:2" json:"age"`
	Gender int    `gorm:"size:2" json:"gender"` //1:男、2:女
	Pwd    string `gorm:"size:255" json:"pwd"`
}

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
