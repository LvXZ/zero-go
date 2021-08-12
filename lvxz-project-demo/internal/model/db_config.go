package model

// @Author: lvxiaozheng
// @Date: 2021/8/12 16:37
// @Description: 数据实体结构绑定数据源表名

func (NotifyMessage) TableName() string {
	return "notify_message"
}

func (UserToken) TableName() string {
	return "user_token"
}

//func AAAa(){
//
//	dd := mapper.Dao2{}
//
//
//	dd.AA = func(userId int64) (userToken UserToken) {
//		return UserToken{}
//	}
//
//	dd.AA
//}
