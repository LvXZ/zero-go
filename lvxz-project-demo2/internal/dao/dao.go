package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 17:10
// @Description:

//全局DB连接（单例）
var DB *gorm.DB

// dao 所有底层数据源配置
type dao struct {
}

// Dao dao interface
type Dao interface {
	Init()

	Close()

	Ping(ctx context.Context) (err error)
}

// Init initial the resource.
func (d *dao) Init() {

}

// Close close the resource.
func (d *dao) Close() {

}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
