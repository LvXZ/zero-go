package utils

import (
	"encoding/json"
	"reflect"
)

// @Author: lvxiaozheng
// @Date: 2021/7/12 14:47
// @Description: 类对象处理

/**
 * 反射Struct Copy
 * 注意成员属性struct指针地址相同
 */
func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		//这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
		dvalue.Set(value)
	}
}

/**
 * JSON转录 Struct Copy
 * 注意成员属性struct指针地址不同
 */
func JSONCopyStruct(src, dst interface{}) {
	srcJSON, _ := json.Marshal(src)
	_ = json.Unmarshal(srcJSON, dst)
}
