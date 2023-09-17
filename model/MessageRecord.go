package model

type MessageRecord struct{
	Id       int `form:"id" gorm:"primaryKey"`
	UserId int  `form:"user_id" gorm:"comment:创建人id;default:0"`
	Type int  `form:"type" gorm:"comment:类型 1-发送者 2-接收者default:1"`
	ListId int  `form:"list_id" gorm:"comment:列表id;default:0"`
	MessageId int  `form:"message_id" gorm:"comment:消息id;default:0"`
	IsRead  int `form:"is_read" gorm:"comment:群聊id;default:0"`
	IsDelete int `form:"is_delete" gorm:"comment:是否已读 1-已读 2-未读;default:1"`
	Created_at int `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&MessageRecord{})
//}