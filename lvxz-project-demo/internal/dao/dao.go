package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 17:10
// @Description:

// dao 所有底层数据源配置
type dao struct {
	db *gorm.DB //mysql数据源
}

// Dao dao interface
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
}

// Close close the resource.
func (d *dao) Close() {

}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
