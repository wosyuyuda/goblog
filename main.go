<<<<<<< HEAD
package main

/*
 * @Description:
 * @Author: longfei
 * @Date: 2021-04-10 09:26:45
 * @LastEditTime: 2021-04-25 17:44:17
 * @LastEditors: longfei
 * @FilePath: \go\main.go
 */

import (
	con "test/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	//先以正常模式运行，下面去掉为调试模式。。
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("view/*")      //这里是引入模板文件
	router.Static("/static", "static") //引入静态目录
	router.GET("/", con.Index)         //首页进入控制器
	router.GET("/aa", con.Index)
	router.Run(":8000") //开启端口访问
}
=======
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
		v1.POST("/sub", con.Login)  //用户登陆提交的接口
		v1.GET("/", con.AdminIndex) //管理页，现在是啥也还没有
	}

	r.Run(":8000") //开启端口访问
}
>>>>>>> caeb3c84e7cad0ff5e4db2262258207fedf886b1
