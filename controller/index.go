package controller

/*
 * @Description:
 * @Author: longfei
 * @Date: 2021-04-22 13:56:46
 * @LastEditTime: 2021-04-25 17:43:30
 * @LastEditors: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"net/http"

	d "test/model"

	"github.com/gin-gonic/gin"
)

func View() {
	d.Views()
}

func Index(c *gin.Context) {
	type student struct {
		Name string
		Age  int8
	}
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	c.HTML(http.StatusOK, "tt.html", gin.H{
		"title":  "Gin",
		"stuArr": [2]*student{stu1, stu2},
	})
}
