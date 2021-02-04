package main

// @Author: lvxiaozheng
// @Date: 2021/2/1 16:48
// @Description: 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// User 用户信息
type User struct {
	Id int `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
	Name string
	Gender string
	Hobby string
}

// DBConfig 数据库yaml配置
type Database struct {
	Datasource 	string	`yaml:"datasource"`
	Url 		string 	`yaml:"url"`
	Username 	string	`yaml:"username"`
	Password 	string	`yaml:"password"`
}


type Application struct {
	Port string
}


type Conf struct {
	Database Database
	Application Application
}

func GetConf() Conf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}

	fmt.Println(string(yamlFile))

	var conf Conf
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return conf
}

func main() {

	dbConfig := Database{}
	dbConfig = GetConf().Database

	db, err := gorm.Open(dbConfig.Datasource, dbConfig.Username+":"+dbConfig.Password+"@"+dbConfig.Url)
	if err!= nil{
		panic(err.Error())
	}
	defer db.Close()

	// 自动迁移
	//db.AutoMigrate(&User{})
	db.SingularTable(true)

	u1 := User{ Name:"lvxz", Gender:"男", Hobby:"篮球"}
	// 创建记录
	db.Create(&u1)
	// 查询

	var uu User
	db.Find(&uu, "hobby=?", "篮球1")
	fmt.Printf("%#v\n", uu)
	// 更新
	db.Model(&uu).Update("hobby", "双色球")
	// 删除
	//db.Delete(&u)
}
