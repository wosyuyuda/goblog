package install

import (
	"fmt"
	"goblog/dao"
	d "goblog/model"
	"goblog/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Install(c *gin.Context) {
	db := dao.MDB
	Idb(db)             //初始化数据库表
	b := Auser(db)      //添加第一个用户,如果已经有用户,那么会返回错误
	l := Atype(db)      //添加第一个分类
	o := Aview(db)      //添加第一篇文章
	g := initConfig(db) //添加网站的标题关键词描述
	if b && l && o && g {
		c.JSON(200, gin.H{"msg": "初始化成功", "code": 200})
	} else {
		c.JSON(200, gin.H{"msg": "初始化失败", "code": 400})
	}
}

//初始化标题描述关键词,后期考虑用前端传过来的数据处理,直接放到配置文件了..好像没啥用了
func initConfig(db *gorm.DB) bool {
	return true
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
		userNew.Pwd = util.Md5jiayan("123456") //md5加盐后加密一下
		return db.Create(&userNew).Error == nil
	} else {
		fmt.Println("后台有用户")
		return false
	}
}

//添加一个分类
func Atype(db *gorm.DB) bool {
	tp := new(d.Tp)
	db.First(&tp)
	if tp.ID != 0 {
		return false
	}
	tp.Name = "简介"
	tp.Info = "第一个分类"
	return db.Create(&tp).Error == nil
}

//添加第一篇文章
func Aview(db *gorm.DB) bool {
	view := new(d.View)
	db.First(&view)
	if view.ID != 0 {
		return false
	}
	view.Title = "欢迎使用goblog"
	view.Typeid = 1
	view.Content = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg 作者原来是个phper转golang了,此代码适合同样转golang的初学者"
	view.Body = "欢迎使用goblog,更多内容可以进入我的gitee https://gitee.com/wosylf/gomybolg 作者原来是个phper转golang了,此代码适合同样转golang的初学者"
	return db.Create(&view).Error == nil
}

//初始化数据库表文件
func Idb(db *gorm.DB) {
	db.AutoMigrate(&d.User{})
	db.AutoMigrate(&d.Tp{})
	db.AutoMigrate(&d.View{})
	db.AutoMigrate(&d.Config{})
}
