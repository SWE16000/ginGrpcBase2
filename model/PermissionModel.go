package model

import initDBGorm "ginGrpcBase2/database/gorm"

type Permission struct {
	Id       int64  `form:"id" gorm:"primaryKey"`
	Name     string `form:"name" gorm:"comment:权限名称;size:255"`
	Componet string `form:"componet" gorm:"权限组件别名,前端import时的名称;size:255"`
	ParentId int64  `form:"parent_id" gorm:"comment:父级id;default:0"`
	Type int64 `form:"type" gorm:"comment:权限类型1.页面2.菜单3.按钮;default:1"`
	Path string `form:"path" gorm:"comment:前端路由地址;size:255;uniqueIndex:path_permission_type"`
	ApiPath string `form:"api_path" gorm:"comment:后端接口地址;size:255"`
	SortId int64 `form:"sort_id" gorm:"comment:排序;default:0"`
	Icon  string `form:"icon" gorm:"comment:权限图标;size:255"`
	IsLock int64 `form:"is_lock" gorm:"comment:是否开启,1开启 2未开启;default:1"`
	IsWhite int64 `form:"is_white" gorm:"comment:是否是白名单,后台不能验证接口权限,1是 2不是;default:2"`
	PermissionType int `form:"permission_type" gorm:"comment:端口分类:1.后台用户 2.服务商用户 3.商户用户;default:1;uniqueIndex:path_permission_type"`
	Created_at int64 `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int64 `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&Permission{})
//}
//查询角色的权限
func PermissionByRole(roleId int64) (res []Permission,err error)  {
	var permissions []Permission
	err=initDBGorm.Db.Model(Permission{}).
		Joins(" left join role_permissions as b on b.permission_id=permissions.id ").
		Where(" b.role_id =? ",roleId).Find(&permissions).Error
	if err!=nil{
		return permissions,err
	}
	return permissions,nil
}
//查询角色权限
func PermissionInRole(roleId []int64)(res []Permission,err error)   {
	var permissions []Permission
	err=initDBGorm.Db.Model(Permission{}).
		Joins(" left join role_permissions as b on b.permission_id=permissions.id ").
		Where(" b.role_id in ?",roleId).Find(&permissions).Error
	if err!=nil{
		return permissions,err
	}
	return permissions,nil

}
func PermissionInRoleWithApiParh(roleId []int64,apiPath string)(res []Permission,err error)   {
	var permissions []Permission
	err=initDBGorm.Db.Model(Permission{}).
		Joins(" left join role_permissions as b on b.permission_id=permissions.id ").
		Where(" b.role_id in ? and permissions.api_path =?",roleId,apiPath).Find(&permissions).Error
	if err!=nil{
		return permissions,err
	}
	return permissions,nil

}