package util

import "unsafe"

type Pages struct {
	Upage    int
	Page     int
	Nextpage int
	Num      int
}

func GetPage(num int64, p int) Pages {
	page := Pages{}

	idPointer := (*int)(unsafe.Pointer(&num))
	znum := *idPointer

	page.Num = znum
	page.Page = p //当前页
	//获取上一页
	if p <= 1 {
		page.Upage = 0
	} else {
		page.Upage = p - 1
	}

	if p*10 > znum {
		page.Nextpage = 0
	} else {
		page.Nextpage = p + 1
	}

	return page
}
