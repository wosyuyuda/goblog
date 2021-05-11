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

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//这个后面加到一个中间件里面去。现在就先用着吧
func Islogin(c *gin.Context) {
	session := sessions.Default(c)

	uid := session.Get("uid")
	fmt.Printf("aa%+v\n", uid)
	//手动删除了cookie后获取到的缓存是nil,没有找到用户信息的时候存的是0
	if uid == 0 || uid == "" || uid == nil { //跳转到登陆界面，这里后面应该放到中间件里面去
		c.Redirect(http.StatusMovedPermanently, "/admin/login")
		return
	}
}

//后台的首页，再在其它页面添加一个文章列表的功能
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_index.html", gin.H{}) //进入管理首页
}

//后台的文章的管理页面
func AdminList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pagenum, _ := strconv.Atoi(page)
	db := d.LinkDb() //连接数据库模型
	v := new(view)
	db.Limit(10).Offset(pagenum).Find(&v)
	//这里模板整一下
	c.HTML(http.StatusOK, "tt.html", gin.H{
		"list": v,
	})
}

//后台的添加文章详情页
func AdminAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_AddView.html", gin.H{}) //进入管理首页
}

//这里加一个接收前端数据的再返回数据就好啦，应该再加一个是否登陆判断
func AddView(c *gin.Context) {
	view1 := new(view)
	type1 := c.PostForm("typeid") //这里分类还要转成int类型，真麻烦,好像直接用string还方便些
	viewType, _ := strconv.Atoi(type1)
	view1.Typeid = viewType

	view1.Title = c.PostForm("title")
	view1.Body = c.PostForm("body")
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
	}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	conn := d.GetDb()
	u := new(User)
	conn.Where("name = ?", name).Find(&u)
	/* fmt.Printf("用户信息是%+v\n", u)
	fmt.Printf("传过来的账号是：%s 密码是：%s", name, pwd)
	fmt.Printf("后台的账号是：%s 密码是：%s，uid是 %d", u.Name, u.Pwd, u.Id) */

	if util.Md5(pwd) != u.Pwd || u.Id == 0 {
		c.JSON(200, gin.H{"msg": "账号不存在或者密码错误", "code": 400})
		c.Abort()
	}
	session := sessions.Default(c)
	option := sessions.Options{MaxAge: 3600 * 8} //8小时后过期
	session.Options(option)
	session.Set("uid", u.Id) //把用户ID存进session,后面拿出来确认
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/admin/")
	c.JSON(200, gin.H{"msg": "登陆成功", "code": 200})

}
