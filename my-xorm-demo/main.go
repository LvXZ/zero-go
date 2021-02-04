package main

// @Author: lvxiaozheng
// @Date: 2021/2/1 18:50
// @Description: 

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// User 用户信息
type User struct {
	Id int `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
	Name string
	Gender string
	Hobby string
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:Qwer1234@@(81.68.89.3:3306)/robot?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}

	var user User
	engine.Where("hobby = '篮球1'").Get(&user)
	content, err := json.Marshal(user)
	fmt.Println(string(content))


}
