package util

import (
	"fmt"
	"strings"
)

func F判断是否为手机访问(userAgent string) bool {
	userAgent = strings.ToLower(userAgent)
	fmt.Println("获取到的useragent", userAgent)
	if len(userAgent) == 0 {
		return false
	}
	mbo := false
	mobileKeywords := []string{"mobile", "android", "silk/", "kindle",
		"blackBerry", "opera mini", "opera mobi"}

	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, mobileKeywords[i]) {
			mbo = true
			break
		}
	}

	return mbo
}
