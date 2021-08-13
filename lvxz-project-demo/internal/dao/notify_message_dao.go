package dao

import (
	"lvxz-project-demo/internal/model"
)

// @Author: lvxiaozheng
// @Date: 2021/8/12 19:32
// @Description: notify_message数据表dao

type NotifyMessageDao interface {
	SelectById(userId int64) (notifyMessage *model.NotifyMessage)
}

func (d *dao) SelectById(id int64) (notifyMessage *model.NotifyMessage) {
	notifyMessage = &model.NotifyMessage{}
	d.db.Where("id = ?", id).Find(&notifyMessage)
	return notifyMessage
}
