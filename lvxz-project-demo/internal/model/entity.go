package model

import "time"

// @Author: lvxiaozheng
// @Date: 2021/8/12 16:24
// @Description: 数据源底层实体结构

//通知消息实体结构
type NotifyMessage struct {
	ID         int64     `json:"id"`
	Content    string    `sql:"content" json:"content"`
	Type       int       `sql:"type" json:"type"`
	Address    string    `sql:"address" json:"address"`
	Status     int       `sql:"status" json:"status"`
	CreateTime time.Time `gorm:"-" json:"-"`
}

//用户Token实体结构
type UserToken struct {
	ID         int64     `json:"id"`
	UserId     int64     `sql:"user_id" json:"userId"`
	Token      string    `sql:"token" json:"token"`
	CreateTime time.Time `gorm:"-" json:"-"`
}
