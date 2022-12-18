package server

import (
	"bytes"
	"fmt"
	"goblog/config"
	"goblog/util"
	"html/template"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//方便后期整一个模板的配置功能,如果是手机端访问，那么获取手机的模板
func F自己写的模板方法(c *gin.Context, temp string, db interface{}) {
	//此处可以加一个获取配置的模板目录方法，这样就能更智能些
	tempdir := config.Configv.GetString("tempdir")
	/* temp = "temp/xing/" + temp */
	dir := "temp" + tempdir + "/"
	userAgent := c.GetHeader("User-Agent") + c.GetHeader("user-agent")
	fmt.Println("ismobile", config.IsMobile)
	//判断当前环境是否为手机，判断是否开户了手机端支持。。
	if config.IsMobile == 1 && util.F判断是否为手机访问(userAgent) {
		dir += "m/"
	}
	content, err := ioutil.ReadFile(dir + temp)
	if err != nil {
		RData(err, c)
		return
	}
	//fmt.Println("body", temp)
	tmpl, err := template.New("tmpl").Delims("{[", "]}").Parse(string(content))
	if err != nil {
		RData(err, c)
		return
	}
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, db)
	if err != nil {
		RData(err, c)
		return
	}
	body := tmplBytes.String()
	F返回字符串(c, body)
}

func F返回字符串(c *gin.Context, str string) {
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(str))
}
