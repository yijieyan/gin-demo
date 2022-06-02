package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

//EncodeByMd5 md5加密
func EncodeByMd5(str string) string {
	w:= md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x",w.Sum(nil))
	return md5Str
}
