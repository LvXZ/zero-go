package mapper

import (
	"lvxz-project-demo2/internal/dao"
	"lvxz-project-demo2/internal/model"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 19:27
// @Description: user_token数据表CURD操作

func SelectByUserId(userId int64) (userToken *model.UserToken) {
	userToken = &model.UserToken{}
	dao.DB.Where("user_id = ?", userId).Find(&userToken)
	return userToken
}
