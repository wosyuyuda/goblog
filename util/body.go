package util

import "regexp"

//去掉内容中的html标签
func HanderHmtl(s string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return re.ReplaceAllString(s, "\n")
}
