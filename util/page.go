package util

import (
	"goblog/model"
	"math"
	"strconv"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type Pages struct {
	Upage    int //上一页
	Page     int //当前页
	Nextpage int //下一页
	Num      int //总共的条数
	Pages    int //总共多少页
}

func PageNum(c *gin.Context) int {
	page := c.DefaultQuery("page", "1")
	pagenum, _ := strconv.Atoi(page)
	return pagenum
}

//初始化列表信息
func Ini(p *model.PageList) {
	//先获取页码
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Num == 0 {
		p.Num = 10
	}
}

//把列表信息处理一下
func PagesinfoTo(p *model.ListInfo) {
	v := p.Views
	for i := 0; i < len(v); i++ {
		if len(v[i].Pic) == 0 {
			v[i].Pic = "/static/no-images.jpg"
		}
		tm := time.Unix(int64(v[i].CreatedAt), 0)
		v[i].Ctime = tm.Format("2006-01-02 15:04:05")
	}
	//处理一下上一页下一页
	page := p.Page
	if page.Page == 0 {
		page.Page = 1
	}
	page.SumPage = int(math.Ceil(float64(page.Sum)/10 + 0.5))
	page.Nex = page.Page + 1
	if page.Nex > page.SumPage {
		page.Nex = 0
	}
	page.Up = page.Page - 1
	if page.Nex < 1 {
		page.Nex = 0
	}
	p.Page = page
}

//num一共多少条,p当前第几页
func GetPage(num int64, p int) Pages {
	page := Pages{}
	off := 10                                 //一页默认10条
	idPointer := (*int)(unsafe.Pointer(&num)) //int64转int
	znum := *idPointer
	page.Pages = int(math.Ceil(float64(znum) / 10)) //看一共有多少页

	page.Num = znum
	page.Page = p //当前页
	//获取上一页
	if p <= 1 {
		page.Upage = 0
	} else {
		page.Upage = p - 1
	}

	if p*off > znum {
		page.Nextpage = 0
	} else {
		page.Nextpage = p + 1
	}

	return page
}
