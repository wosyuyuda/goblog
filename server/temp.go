package server

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//方便后期整一个模板的配置功能,
func F自己写的模板方法(c *gin.Context, temp string, db interface{}) {
	//此处可以加一个获取配置的模板目录方法，这样就能更智能些
	temp = "temp/xing/" + temp
	content, err := ioutil.ReadFile(temp)
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
