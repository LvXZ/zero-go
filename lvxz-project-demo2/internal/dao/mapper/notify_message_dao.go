package mapper

import (
	"lvxz-project-demo2/internal/dao"
	"lvxz-project-demo2/internal/model"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 19:32
// @Description: notify_message数据表CURD操作

func SelectById(id int64) (notifyMessage *model.NotifyMessage) {
	notifyMessage = &model.NotifyMessage{}
	dao.DB.Where("id = ?", id).Find(&notifyMessage)
	return notifyMessage
}
