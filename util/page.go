package util

import (
	"math"
	"unsafe"
)

type Pages struct {
	Upage    int //上一页
	Page     int //当前页
	Nextpage int //下一页
	Num      int //总共的条数
	Pages    int //总共多少页
}

func GetPage(num int64, p int) Pages {
	page := Pages{}
	off := 10 //一页默认10条
	idPointer := (*int)(unsafe.Pointer(&num))
	znum := *idPointer

	page.Pages = int(math.Floor(float64(znum/10))) + 1 //看一共有多少页

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
