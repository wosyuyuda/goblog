package util

import "strings"

func F判断是否为手机访问(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}
	isMobile := false
	mobileKeywords := []string{"Mobile", "Android", "Silk/", "Kindle",
		"BlackBerry", "Opera Mini", "Opera Mobi"}

	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, mobileKeywords[i]) {
			isMobile = true
			break
		}
	}

	return isMobile
}
