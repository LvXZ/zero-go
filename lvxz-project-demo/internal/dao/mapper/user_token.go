package mapper

import (
	"gorm.io/gorm"
	"lvxz-project-demo/internal/model"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 16:51
// @Description: 数据表CURD操作

//通过userId查询
func SelectByUserId(db *gorm.DB, userId int64) (userToken *model.UserToken) {
	userToken = &model.UserToken{}
	db.Where("user_id = ?", userId).Find(&userToken)
	return userToken
}

func SelectByUserId2(db *gorm.DB, userId int64) (userToken model.UserToken) {
	db.Where("user_id = ?", userId).Find(&userToken)
	return userToken
}
