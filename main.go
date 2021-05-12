package main

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\main.go
 */

import (
	"net/http"
	con "test/controller"

	"test/middleware"
	"test/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//先以正常模式运行，下面去掉为调试模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.LoadHTMLGlob("view/*")        //这里是引入模板文件
	r.Static("/static", "static")   //引入静态目录
	r.GET("/", con.Lists1)          //这个是首页，模板整整就可以啦
	r.GET("/list/:id", con.Views)   //具体列表页
	r.GET("/view/:id", con.GetView) //文章详情页，这里的详情页可以开始获取数据了

	//设置session开始让下面调用
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//这里加一个判断是否登陆的中间件，如果没有缓存的用户ID，直接跳出到登陆页面

	v2 := r.Group("/admin")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login1.html", gin.H{})
		}) //登陆页
		v2.POST("/sub", con.Login) //用户登陆提交的接口
	}

	r.Use(middleware.Islogin)

	v1 := r.Group("/admin")
	{
		v1.GET("/", con.AdminIndex)    //管理页，现在是啥也还没有
		v1.GET("/list", con.AdminList) //后台的文章列表，这里要加一个管理选项
		//v1.GET("/addView1", con.AdminAddView) //添加文章界面，笑死，gin模板语法跟vue.js模板语法冲突，已转到静态页面'/static/view/admin_addview.html'
		v1.GET("/gettype", con.Gt) //后台的文章列表，这里要加一个管理选项
		v1.Any("/ueditor/controller", util.Action)

		v1.POST("/addView", con.AddView)  //添加文章
		v1.POST("/addUser", con.AddU)     //添加用户
		v1.POST("/addType", con.AddTypes) //添加分类
	}

	r.Run(":8000") //开启端口访问,本地再试一下提交
}
