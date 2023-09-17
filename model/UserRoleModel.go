package model

type UserRole struct{
	Id int64 `form:"id" gorm:"primaryKey"`
	RoleId int64 `form:"role_id" gorm:"comment:角色id;uniqueIndex:user_role_index"`
	UserId int64 `form:"user_id" gorm:"comment:用户id;uniqueIndex:user_role_index"`
	IsDelete int64 `from:"is_delete" gorm:"comment:是否删除 1-未删除 2-已删除;default:1"`
	Created_at int64 `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int64 `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	////数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&UserRole{})
//	////添加默认数据
//	initDBGorm.Db.Create(&UserRole{Id: 1,RoleId:1,UserId: 1})
//}
