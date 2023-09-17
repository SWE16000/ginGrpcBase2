package model

type RolePermission struct{
	Id int64 `form:"id" json:"id" gorm:"primaryKey"`
	RoleId int64 `form:"role_id" json:"role_id"  gorm:"uniqueIndex:role_permission"`
	PermissionId int64 `form:"permission_id" json:"permission_id" gorm:"uniqueIndex:role_permission"`
	Created_at int64 `form:"created_at" json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int64 `form:"updated_at" json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}

//func init()  {
//	initDBGorm.Db.AutoMigrate(&RolePermission{})
//}