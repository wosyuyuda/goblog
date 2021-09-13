package model

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\model\model.go
 */

import (
	"gorm.io/gorm"
)

//下面直接转过来不行
type View struct {
	ID        uint           `gorm:"primarykey"  json:"id" form:"id"`
	CreatedAt uint           `gorm:"size:255" json:"createtime" form:"createtime"`
	UpdatedAt uint           `gorm:"size:255" json:"updatetime" form:"updatetime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Typeid    int            `gorm:"size:10"  json:"typeid" form:"typeid"`          //分类的ID，关联
	Title     string         `gorm:"size:255" json:"title" form:"title"`            //标题
	Click     int            `gorm:"size:10" json:"click"`                          //点击量
	Tuijian   int8           `gorm:"size:1" json:"tuijian" form:"tuijian"`          //是否为推荐
	Swiper    int8           `gorm:"size:1" json:"swiper" form:"swiper"`            //是否为轮播图
	Pic       string         `gorm:"size:255" json:"pic" form:"pic"`                //文章的缩略图
	Status    int8           `gorm:"size:1;default:1;" json:"status" form:"status"` //文章状态，0删除，1正常
	Content   string         `gorm:"size:500" json:"content" form:"content"`        //文章的简介
	Body      string         `gorm:"size:5000" json:"body" form:"body"`             //详细的内容
	Ctime     string         `gorm:"-" json:"ctime" form:"ctime"`
	Typename  string         `gorm:"-" json:"typename" form:"typename"`
	//Tps      Tp     `json:"tps" gorm:"-"`                           //这里放分类信息types
}

//评论
type Comment struct {
	ID        uint   `gorm:"primarykey"  json:"id" form:"id"`
	CreatedAt uint   `gorm:"size:255" json:"createtime" form:"createtime"`
	Status    int8   `gorm:"size:1;default:1;" json:"status" form:"status"` //状态，0删除，1正常
	Uid       uint   `gorm:"index"  json:"uid" form:"uid"`                  //发表评论的ID
	Name      string `gorm:"size:24" json:"name" form:"name"`               //发表评论的id
}

type ViewJson struct {
	View
	ID        uint `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt uint
	Typeid    int    ` json:"typeid" form:"typeid"`     //分类的ID，关联
	Title     string ` json:"title" form:"title"`       //标题
	Pic       string ` json:"pic" form:"pic"`           //文章的缩略图
	Status    uint8  ` json:"status"`                   //文章状态，0删除，1正常
	Content   string `json:"content" form:"content"`    //文章的简介
	Typename  string ` json:"typename" form:"typename"` //存储一下分类名称
	Ctime     string ` json:"ctime" form:"ctime"`
}

//分类表
type Tp struct {
	ID     uint       `gorm:"primarykey" json:"id" form:"id"`
	Name   string     `gorm:"size:255;" form:"name" json:"name"`               //名称
	Level  int        `gorm:"size:10;default:0;" form:"level" json:"level"`    //分类排序
	Info   string     `gorm:"size:255;" form:"info" json:"info"`               //介绍
	Status string     `gorm:"size:1;default:'1';" form:"status" json:"status"` //状态,默认1,0表示删除
	Views  []ViewJson `gorm:"-"`
}

//系统设置
type Config struct {
	ID      uint   `gorm:"primarykey"  json:"id" form:"id"`        //设置ID
	Name    string `gorm:"size:255"  json:"name" form:"name"`      //设置的名称
	Type    string `gorm:"size:10"  json:"type" form:"type"`       //设置所属分类
	Content string `gorm:"size:500" json:"content" form:"content"` //设置的介绍
	Value   string `gorm:"size:500" json:"value" form:"value"`     //系统设置值
	Group   uint   `gorm:"size:6" json:"group" form:"group"`       //所属群组,方便一次调出来
}

//网站的标题关键词等.
type Tdk struct {
	Title       string ` json:"title" form:"title"`
	Keyword     string ` json:"keyword" form:"keyword"`
	Description string ` json:"description" form:"description"`
	Beian       string ` json:"beian" form:"beian"`
	Goan        string ` json:"goan" form:"goan"`         //公安的备案号
	Goanurl     string ` json:"goanurl" form:"goanurl"`   //备案的网址
	Tongjiid    string ` json:"tongjiid" form:"tongjiid"` //百度统计的id
}

//用户表
type User struct {
	Id     int    `gorm:"primary_key" json:"id"  form:"id"`
	Name   string `gorm:"size:255" json:"name"  form:"name"`
	Age    int    `gorm:"size:2" json:"age" form:"age"`
	Gender int    `gorm:"size:2" json:"gender" form:"gender"` //1:男、2:女
	Pwd    string `gorm:"size:255" json:"pwd" form:"pwd"`
}
