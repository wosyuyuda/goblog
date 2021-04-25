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
	//先以正常模式运行，下面去掉为调试模式
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("view/*")      //这里是引入模板文件
	router.Static("/static", "static") //引入静态目录
	router.GET("/", con.Index)         //首页进入控制器
	router.Run(":8000")                //开启端口访问
}
