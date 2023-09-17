package model

type MessageGroupUser struct {
	Id       int  `form:"id" gorm:"primaryKey"`
	MessageGroupId  int  `form:"message_group_id" gorm:"comment:群聊id;default:0"`
	UserId int  `form:"user_id" gorm:"comment:创建人id;default:0"`
	IsDelete int  `form:"is_delete" gorm:"comment:是否删除 1-未删除 2-已删除;default:1"`
	Created_at int `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&MessageGroupUser{})
//}