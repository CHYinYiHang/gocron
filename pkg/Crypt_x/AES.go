package Aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//AES加密算法
//采用CBC模式+PKCS7 填充方式实现AES的加密和解密
//yinyihang 2020/03/28
//单元测试：数据为100字节
//BenchmarkAES-6     	 1836958	       655 ns/op
//BenchmarkAES-6     	 1820587	       652 ns/op
//BenchmarkAES-6     	 1843641	       654 ns/op
//BenchmarkAES-6     	 1834952	       651 ns/op
//BenchmarkAES-6     	 1836258	       658 ns/op

//补码
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//加密数据
//参数：@1原数据 @2密钥，类型：[]byte 字节数组
func EncryptAES(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
//参数：@1密文 @2加密时使用的密钥，类型：[]byte 字节数组
func DecryptAES(nowData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(nowData))
	blockMode.CryptBlocks(origData, nowData)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
