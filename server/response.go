package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 403
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
	c.Abort()
}

//直接处理最后的结果，然后这里判断并且给出答复成功或者失败
func IfRes(err error, c *gin.Context) {
	if err != nil {
		//如果失败直接返回错误
		FailWithMessage(err.Error(), c)
	} else {
		//否则返回失败
		Ok(c)
	}
}

//返回错误判断，如果不等于nil那么把错误信息返回，否则返回正确的数据
func ResDataError(data interface{}, err error, c *gin.Context) {
	if err != nil {
		//如果失败直接返回错误
		FailWithMessage(err.Error(), c)
	} else {
		//否则返回成功
		OkWithData(data, c)
	}
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

//返回OK，并且提示成功
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

//返回数据并提示成功
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

//返回OK，并且把数据跟提示返回
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

//返回失败
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

//返回操作失败的提示
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

//返回错误7，并提示信息，还有返回的数据
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
