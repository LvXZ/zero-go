package service

import (
	"context"
	"fmt"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 19:41
// @Description:

func (service *Service) QueryBaseInfo(ctx context.Context, userId int64) {
	userToken := service.dao.SelectByUserId3(userId)
	fmt.Println(userToken)
}
