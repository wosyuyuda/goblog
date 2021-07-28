package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

const ueBasePath = `static/upload/`

func Action(c *gin.Context) {
	action := c.Query("action")
	datePath := time.Now().Format("20060102") + `/`

	switch action {
	//自动读入配置文件，只要初始化UEditor即会发生
	case "config":
		jsonByte, _ := ioutil.ReadFile("static/ueditor/conf/config.json")
		re, _ := regexp.Compile(`\\/\\*[\\S\\s]+?\\*\\/`)
		jsonByte = re.ReplaceAll(jsonByte, []byte(""))
		c.Writer.Write(jsonByte)

	case "uploadimage":
		{
			//创建保存文件的目录，每天一个目录
			err := os.MkdirAll(ueBasePath+`images/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, header, err := c.Request.FormFile("upfile")

			file1, header1, err1 := c.Request.FormFile("file") //这里兼容普通的图片上传
			if err != nil {
				if err1 == nil {
					file = file1
					header = header1
				} else {
					data, _ := json.Marshal(map[string]string{
						"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
					})
					c.Writer.Write(data)
					return
				}
			}
			var Suffix = regexp.MustCompile(`\.+[a-z]+`)
			name := fmt.Sprintf("%d", time.Now().UnixNano()/1e6) + Suffix.FindString(header.Filename)
			out, err := os.Create(ueBasePath + `images/` + datePath + name)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			data, err := json.Marshal(map[string]string{
				"url":      fmt.Sprintf("/%s", ueBasePath+`images/`+datePath+name), //保存后的文件路径
				"title":    "",                                                     //文件描述，对图片来说在前端会添加到title属性上
				"original": header.Filename,                                        //原始文件名
				"state":    "SUCCESS",                                              //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			if err != nil {
				panic(err)
			}
			c.Writer.Write(data)
		}

	case "uploadvideo":
		{
			err := os.MkdirAll(ueBasePath+`video/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的视频
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, err := c.FormFile("upfile")
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			name := time.Now().Format("150405") + file.Filename
			path := ueBasePath + `video/` + datePath
			if err := c.SaveUploadedFile(file, path+name); err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"url":      fmt.Sprintf(ueBasePath + `images/` + datePath + name), //保存后的文件路径
				"title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
				"original": file.Filename,                                         //原始文件名
				"state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			c.Writer.Write(data)
		}

	case "uploadfile":
		{
			//创建保存文件的目录，每天一个目录
			err := os.MkdirAll(ueBasePath+`files/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的文件
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, header, err := c.Request.FormFile("upfile")
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			name := time.Now().Format("150405") + header.Filename
			out, err := os.Create(ueBasePath + `files/` + datePath + name)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"url":      fmt.Sprintf(ueBasePath + `images/` + datePath + name), //保存后的文件路径
				"title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
				"original": header.Filename,                                       //原始文件名
				"state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			c.Writer.Write(data)
		}
	case "uploadscrawl":
		{
			path := ueBasePath + `scrawl/` + datePath
			name := time.Now().Format("150405.999999") + `.jpg`
			err := os.MkdirAll(path, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//upfile为base64格式文件，转成图片保存
			upfile, _ := c.GetPostForm("upfile")
			upBytes, err := base64.StdEncoding.DecodeString(upfile) // + "_" + filename
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取图片错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			err = ioutil.WriteFile(path+name, upBytes, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("保存涂鸦文件错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"state":    "SUCCESS",
				"url":      `/` + path + name,
				"title":    `涂鸦`,
				"original": `涂鸦不见了`,
			})
			c.Writer.Write(data)
			return
		}
	}
}
