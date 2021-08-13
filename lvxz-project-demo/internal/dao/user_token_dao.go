package dao

import (
	"lvxz-project-demo/internal/model"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 19:27
// @Description: user_token数据表dao

type UserTokenDao interface {
	SelectByUserId(userId int64) (userToken *model.UserToken)
}

func (d *dao) SelectByUserId(userId int64) (userToken *model.UserToken) {
	userToken = &model.UserToken{}
	d.db.Where("user_id = ?", userId).Find(&userToken)
	return userToken
}
