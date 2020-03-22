package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5V1(str string) string {
	b := md5.New()
	b.Write([]byte(str))
	return hex.EncodeToString(b.Sum(nil))
}
