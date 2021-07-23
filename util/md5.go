package util

import (
	"crypto/md5"
	"encoding/hex"
)

func tmd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//md5加盐
func Md5jiayan(s string) string {
	return tmd5(s + "a")
}
