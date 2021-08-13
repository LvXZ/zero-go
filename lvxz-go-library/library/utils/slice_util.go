package utils

import "strings"

// @Author: lvxiaozheng
// @Date: 2021/6/21 11:07
// @Description: slice工具类

//集合是否存在目标(区分大小写) true存在 false不存在
func SliceContain(slice []string, str string) bool {
	for _, b := range slice {
		if strings.EqualFold(b, str) {
			return true
		}
	}
	return false
}

//集合是否包含目标(区分大小写) true存在 false不存在
func SliceStringContain(slice []string, str string) bool {
	for _, b := range slice {
		if strings.Contains(str, b) {
			return true
		}
	}
	return false
}
