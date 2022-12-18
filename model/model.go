package model

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \gomybolg\model\model.go
 */

import (
	"gorm.io/gorm"
)

//下面直接转过来不行
type View struct {
	ID        int            `gorm:"primarykey"  json:"id" form:"id"`
	CreatedAt int            `gorm:"size:255" json:"createtime" form:"createtime"`
	UpdatedAt int            `gorm:"size:255" json:"updatetime" form:"updatetime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Typeid    int            `gorm:"size:10"  json:"typeid" form:"typeid"`             //分类的ID，关联
	Click     int            `gorm:"size:10" json:"click"`                             //点击量
	Tuijian   int8           `gorm:"size:1" json:"tuijian" form:"tuijian"`             //是否为推荐
	Swiper    int8           `gorm:"size:1" json:"swiper" form:"swiper"`               //是否为轮播图
	Status    int8           `gorm:"size:1;default:1;" json:"status" form:"status"`    //文章状态，0删除，1正常
	Title     string         `gorm:"type:varchar(255);" json:"title" form:"title"`     //标题
	Appid     string         `gorm:"type:char(16);" json:"appid" form:"appid"`         //appid，如果是推荐小程序的文章
	Pic       string         `gorm:"type:varchar(255);" json:"pic" form:"pic"`         //文章的缩略图
	Content   string         `gorm:"type:varchar(256);" json:"content" form:"content"` //文章的简介
	Body      string         `gorm:"type:longtext;" json:"body" form:"body"`           //详细的内容
	Ctime     string         `gorm:"-" json:"ctime" form:"ctime"`
	Tempdir   string         `gorm:"-" json:"tempdir" form:"tempdir"`
	Typename  string         `gorm:"-" json:"typename" form:"typename"`
	Isapp     bool           `gorm:"-" json:"isapp" form:"isapp"`     //是否为app，如果带了appid，那么就有，就生成二维码
	Comment   []Comment      `gorm:"-" json:"comment" form:"comment"` //评论列表
	//Tps      Tp     `json:"tps" gorm:"-"`                           //这里放分类信息types
}

//评论
type Comment struct {
	ID        int    `gorm:"primarykey"  json:"id" form:"id"`
	CreatedAt int    `gorm:"size:255" json:"createtime" form:"createtime"`
	Status    int8   `gorm:"size:1;default:1;" json:"status" form:"status"` //状态，0删除，1正常,2审核
	Body      string `gorm:"type:varchar(512);" json:"body" form:"body"`    //详细的内容
	Vid       int    `gorm:"index"  json:"vid" form:"vid"`                  //发表评论的ID
	Name      string `gorm:"size:24" json:"name" form:"name"`               //发表评论的id
}

type ViewJson struct {
	View
	CreatedAt int
	ID        int    `gorm:"primarykey" json:"id" form:"id"`
	Typeid    int    ` json:"typeid" form:"typeid"`     //分类的ID，关联
	Title     string ` json:"title" form:"title"`       //标题
	Pic       string ` json:"pic" form:"pic"`           //文章的缩略图
	Status    int8   ` json:"status"`                   //文章状态，0删除，1正常
	Content   string `json:"content" form:"content"`    //文章的简介
	Typename  string ` json:"typename" form:"typename"` //存储一下分类名称
	Ctime     string ` json:"ctime" form:"ctime"`
}

//分类表
type Tp struct {
	ID         int        `gorm:"primarykey" json:"id" form:"id"`
	Name       string     `gorm:"size:255;" form:"name" json:"name"`                       //名称
	Level      int        `gorm:"size:10;default:0;" form:"level" json:"level"`            //分类排序
	Info       string     `gorm:"size:255;" form:"info" json:"info"`                       //介绍
	Status     string     `gorm:"size:1;default:'1';" form:"status" json:"status"`         //状态,默认1,0表示删除
	Tempdir    string     `gorm:"type:varchar(32);" form:"tempdir" json:"tempdir"`         //分类模板目录
	ArcTempdir string     `gorm:"type:varchar(32);" form:"arc_tempdir" json:"arc_tempdir"` //内容模板目录
	SkipPath   string     `gorm:"type:varchar(128);" form:"skip_path" json:"skip_path"`    //跳转的路径
	IsTrue     bool       `gorm:"-"  json:"is_true"`                                       //是否选中
	IsTz       bool       `gorm:"-"  json:"is_tz"`                                         //是否跳转
	Views      []ViewJson `gorm:"-"`
}

//系统设置
type Config struct {
	ID      int    `gorm:"primarykey"  json:"id" form:"id"`        //设置ID
	Name    string `gorm:"size:255"  json:"name" form:"name"`      //设置的名称
	Type    string `gorm:"size:10"  json:"type" form:"type"`       //设置所属分类
	Content string `gorm:"size:500" json:"content" form:"content"` //设置的介绍
	Value   string `gorm:"size:500" json:"value" form:"value"`     //系统设置值
	Group   int    `gorm:"size:6" json:"group" form:"group"`       //所属群组,方便一次调出来
}

//网站的标题关键词等.
type Tdk struct {
	Title       string ` json:"title" form:"title"`
	Keyword     string ` json:"keyword" form:"keyword"`
	Description string ` json:"description" form:"description"`
	Beian       string ` json:"beian" form:"beian"`
	Goan        string ` json:"goan" form:"goan"`         //公安的备案号
	Url         string ` json:"url" form:"url"`           //当前网址,以/结尾
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
