package model

type Message struct{
	Id       int  `form:"id" json:"id" gorm:"primaryKey"`
	Content    string `form:"content" json:"content" gorm:"comment:聊天内容;size:2555"`
	Type       int  `form:"type" json:"type" gorm:"comment:类型:1-文字2-图片3-视频4-音频5-其他文件;default:1"`
	MessageGroupId int  `form:"message_group_id" json:"message_group_id" gorm:"comment:聊天群id;default:0"`
	ListId int  `form:"list_id" json:"list_id" gorm:"comment:列表id;default:0"`
	UserId int  `form:"user_id" json:"user_id" gorm:"comment:发送者id;default:0"`
	State  int  `form:"state" json:"state" gorm:"comment:发送状态 1-成功 2-失败';default:1"`
	Remark  string  `form:"remark" remark:"id" gorm:"comment:发送失败原因';default:''"`
	Created_at int `form:"created_at" json:"created_at"  gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int `form:"updated_at"json:"updated_at"  gorm:"autoUpdateTime;comment:更新时间"`
}


type MessageReturn struct{
	Content    string `form:"name"`
	Type       int  `form:"type"`
	ListId int  `form:"list_id"`
	UserId int  `form:"user_id"`
	WebsocketType string `form:"websocket_type"`
	Status string `form:"status"`
	Message string `form:"message"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&Message{})
//}