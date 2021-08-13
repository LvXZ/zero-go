package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

// encryptMD5 md5加密
func encryptMD5(message string) (result string, err error) {
	if len(message) == 0 {
		err = errors.New("invalid response")
		return
	}

	m := md5.New()
	m.Write([]byte(message))
	result = hex.EncodeToString(m.Sum(nil))
	return
}

// DefaultGetSign 请求获取sign
func DefaultGetSign(appKey string, secret string) (sign string) {

	// 判断输入不出空
	if len(appKey) == 0 || len(secret) == 0 {
		return ""
	}

	// 请求时间戳
	current := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println("请求时间: ", current)

	// 把所有参数名和参数值串在一起
	query := secret + "appKey" + appKey + "timestamp" + current + "version" + "1.0" + secret
	//fmt.Println("请求参数: ", query)

	// md5加密
	md5String, err := encryptMD5(query)
	if err != nil {
		return ""
	}
	//fmt.Println("请求签名: ", md5String)
	return strings.ToUpper(md5String)
}

// GetSign 请求获取sign
func GetSign(appKey string, secret string, timestamp string, version string) (sign string) {

	// 判断输入不出空
	if len(appKey) == 0 || len(secret) == 0 || len(timestamp) == 0 || len(version) == 0 {
		return ""
	}

	// 把所有参数名和参数值串在一起
	query := secret + "appKey" + appKey + "timestamp" + timestamp + "version" + version + secret
	//fmt.Println("请求参数: ", query)

	// md5加密
	md5String, err := encryptMD5(query)
	if err != nil {
		return ""
	}
	//fmt.Println("请求签名: ", md5String)
	return strings.ToUpper(md5String)
}
