package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\admin.go
 */
import (
	"fmt"
	"net/http"
	"strconv"
	d "test/model"
	"test/util"

	"github.com/gin-gonic/gin"
)

//后台的首页，再在其它页面添加一个文章列表的功能
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_index.html", gin.H{}) //进入管理首页
}

type editview struct {
	tp1  []Tp
	view view
}

func AdminGetId(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("进入获取文章")
	db := d.LinkDb()
	l := editview{}
	err := db.Where("id = ?", id).Find(&l.view).Error
	db.Find(&l.tp1)
	fmt.Printf("获取到的内容是%+v", l)
	if err != nil {
		c.JSON(200, gin.H{"msg": "获取失败", "code": 204})
	}

	c.JSON(200, gin.H{"msg": "获取成功", "code": 200, "data": l.view, "type": l.tp1})
}

//后台的文章的管理页面
func AdminList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pagenum, _ := strconv.Atoi(page)
	v := new(view)
	list := v.Findlist("0", pagenum)
	fmt.Println(pagenum)

	//获取全部列表的信息
	var i int64
	db := d.LinkDb()
	db.Model(&view{}).Count(&i)
	p := util.GetPage(i, pagenum)

	fmt.Printf("分页的内容是%+v", p)
	//这里模板整一下
	c.HTML(http.StatusOK, "admin_list.html", gin.H{
		"list":     list,
		"page":     pagenum,
		"pageinfo": p,
	})
}

//后台的添加文章详情页,笑死，gin模板语法跟vue.js模板语法冲突，不能直接用，直接加载静态页面吧
func AdminAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_addview.html", gin.H{}) //进入管理首页
}

//这里加一个接收前端数据的再返回数据就好啦，应该再加一个是否登陆判断
func AddView(c *gin.Context) {
	var data view
	c.ShouldBind(&data)

	//typeid没办法直接拿过来，只好再单独获取了
	typeid := c.PostForm("typeid")
	t, _ := strconv.Atoi(typeid)
	fmt.Printf("%+v", data)
	view1 := new(view)
	view1.Typeid = t
	view1.Title = data.Title
	view1.Body = data.Body
	view1.Swiper = data.Swiper
	view1.Tuijian = data.Tuijian

	//这个截取太麻烦了，直接不用了，由前端去处理数据感觉方便多
	view1.Content = string([]rune(util.HanderHmtl(view1.Body))[:64]) //去掉html标签后截取255的长度放到简介里面去
	fmt.Printf("传过来的标题是：%s 密码是：%s", c.PostForm("body"), c.PostForm("title"))
	//fmt.Println(view1)
	conn := d.GetDb()
	conn.AutoMigrate(&view{})
	err := conn.Create(view1).Error
	if err != nil {
		fmt.Println("创建失败")
	}
	fmt.Println("创建成功")
	c.JSON(200, gin.H{"msg": "创建成功", "code": 200})
}

//用户登陆提交的页面
func Login(c *gin.Context) {
	code := c.PostForm("code")
	if code != "1111" { //此处为验证码验证，后期再扩展
		c.JSON(200, gin.H{"msg": "验证码错误", "code": 400})
		c.Abort()
		return
	}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	conn := d.GetDb()
	u := new(User)
	conn.Where("name = ?", name).Find(&u)
	if util.Md5(pwd) != u.Pwd || u.Id == 0 {
		c.JSON(200, gin.H{"msg": "账号不存在或者密码错误", "code": 400})
		c.Abort()
		return
	}

	util.SetSession(c, "uid", u.Id) //把用户的ID存进session
	c.JSON(200, gin.H{"msg": "登陆成功", "code": 200})

}
