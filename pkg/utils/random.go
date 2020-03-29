package utils

import (
	"math/rand"
	"time"
)

//随机生成自定长度的字符串
//yinyihang 2020/03/29

//性能测试 生成23位字符串
//goos: windows
//goarch: amd64
//pkg: github.com/CHYinYiHang/gocron/test
//BenchmarkBytesMaskImprSrc-6     13829966                85.3 ns/op
//BenchmarkBytesMaskImprSrc-6     13986355                85.6 ns/op
//BenchmarkBytesMaskImprSrc-6     13825648                85.5 ns/op
//BenchmarkBytesMaskImprSrc-6     13830062                85.2 ns/op
//BenchmarkBytesMaskImprSrc-6     13986355                85.4 ns/op
//BenchmarkBytesMaskImprSrc-6     13830364                85.6 ns/op
//FAIL    github.com/CHYinYiHang/gocron/test      8.086s

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImpSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

//参数：@len 渴望的长度
func RandOmStr(len int) string {
	return randStringBytesMaskImpSrc(len)
}
