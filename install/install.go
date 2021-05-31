package install

import (
	"fmt"
	d "goblog/model"
	"goblog/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Install(c *gin.Context) {
	db := d.GetDb()
	Idb(db)             //初始化数据库表
	b := Auser(db)      //添加第一个用户,如果已经有用户,那么会返回错误
	e := Atype(db)      //添加第一个分类
	f := Aview(db)      //添加第一篇文章
	g := initConfig(db) //添加网站的标题关键词描述
	if b && e && f && g {
		c.JSON(200, gin.H{"msg": "初始化成功", "code": 200})
	} else {
		c.JSON(200, gin.H{"msg": "初始化失败", "code": 400})
	}
}

//初始化标题描述关键词,后期考虑用前端传过来的数据处理
func initConfig(db *gorm.DB) bool {
	//设置默认的系统设置
	//先设置tdk标题跟描述,
	var config1 = []d.Config{
		{
			Name:    "title",
			Type:    "string",
			Content: "网站的标题",
			Value:   "goblog",
			Group:   1,
		},
		{
			Name:    "keyword",
			Type:    "string",
			Content: "网站的关键词",
			Value:   "goblog,博客",
			Group:   1,
		},
		{
			Name:    "description",
			Type:    "string",
			Content: "网站的描述",
			Value:   "golang的博客,改这些东西后面再看趴",
			Group:   1,
		},
	}
	//批量创建并返回创建结果
	return db.CreateInBatches(config1, 3).Error == nil
}

//添加用户
func Auser(db *gorm.DB) bool {
	var user d.User
	db.First(&user)
	if user.Id == 0 {
		userNew := new(d.User)
		userNew.Name = "longfei"
		userNew.Pwd = util.Md5("123456") //md5加密一下
		return db.Create(&userNew).Error == nil
	} else {
		fmt.Println("后台有用户")
		return false
	}
}

//添加一个分类
func Atype(db *gorm.DB) bool {
	tp := new(d.Tp)
	tp.Name = "最新"
	tp.Info = "第一个分类"
	return db.Create(&tp).Error == nil
}

//添加第一篇文章
func Aview(db *gorm.DB) bool {
	view := new(d.View)
	view.Title = "这是第goblog的第一篇文章"
	view.Typeid = 1
	view.Content = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg"
	view.Body = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg"
	return db.Create(&view).Error == nil
}

//初始化数据库表文件
func Idb(db *gorm.DB) {
	db.AutoMigrate(&d.User{})
	db.AutoMigrate(&d.Tp{})
	db.AutoMigrate(&d.View{})
	db.AutoMigrate(&d.Config{})
}
