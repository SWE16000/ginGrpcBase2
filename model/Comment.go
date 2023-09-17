package model

type Comment struct{
	Id int `form:"id" json:"id" gorm:"primary"`
	Content string `form:"content" json:"content" gorm:"comment:评论内容;size:2555;default:''"`
	UserId int `form:"user_id" json:"user_id" gorm:"comment:用户id;size:10;default:0"`
	Major string `form:"major" json:"major" gorm:"comment:专业;size:255;default:''"`
	School string `form:"school" json:"school" gorm:"comment:学校;size:255;default:''"`
	Education string `form:"education" json:"education" gorm:"comment:学历;size:255;default:''"`
	Like int `form:"like" json:"like" gorm:"comment:点赞数;size:11;default:0"`
	CreatedAt int `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt int `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&Comment{})
//}