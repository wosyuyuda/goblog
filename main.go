package main

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\main.go
 */

import (
	"net/http"
	con "test/controller"

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
	r.GET("/", con.Lists1)          //这个是首页
	r.GET("/list/:id", con.Views)   //具体列表页
	r.GET("/view/:id", con.GetView) //文章详情页

	//设置session开始让下面调用
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("/admin")
	{

		v1.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login1.html", gin.H{})
		}) //登陆页
		v1.POST("/sub", con.Login)       //用户登陆提交的接口
		v1.GET("/", con.AdminIndex)      //管理页，现在是啥也还没有
		v1.POST("/addView", con.AddView) //添加文章
	}

	r.Run(":8000") //开启端口访问,本地再试一下提交
}
