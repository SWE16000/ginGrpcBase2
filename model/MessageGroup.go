package model

type MessageGroup struct{
	Id       int `form:"id" gorm:"primaryKey"`
	Name    string `form:"name" gorm:"comment:群名称;size:255"`
	Number  int  `form:"number" gorm:"comment:群人数;default:0"`
	Pic    string `form:"pic" gorm:"comment:群头像;size:255;defult:''"`
	UserId int  `form:"user_id" gorm:"comment:创建人id;default:0"`
	Created_at int `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&MessageGroup{})
//}