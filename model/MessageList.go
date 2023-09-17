package model

import (
	"reflect"
)

type MessageList struct {
	Id       int  `form:"id" gorm:"primaryKey"`
	MessageGroupId  int  `form:"message_group_id" gorm:"comment:群聊id;default:0"`
	SendId int  `form:"send_id" gorm:"comment:发送人id;default:0"`
	ReceiveId int  `form:"receive_id" gorm:"comment:接收人id;default:0"`
	Created_at int `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&MessageList{})
//}
func MessageListIsEmpty(messageList MessageList)bool  {
	return reflect.DeepEqual(messageList, MessageList{})
}