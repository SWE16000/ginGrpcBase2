package model

import (
	"fmt"
	initDBGorm "ginGrpcBase2/database/gorm"
	"ginGrpcBase2/utils"
	"time"
)

type Role struct {
	Id   int64  `form:"id" gorm:"primaryKey"`
	Name string `form:"name" gorm:"comment:角色名称;size:255;uniqueIndex:name_permission_type"`
	IsDelete int64 `from:"is_delete" gorm:"comment:是否删除 1-未删除 2-已删除;default:1"`
	PermissionType int64 `form:"permission_type" gorm:"comment:端口分类:1.后台用户 2.服务商用户 3.商户用户;default:1;uniqueIndex:name_permission_type"`
	Created_at int64 `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int64 `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
	//Users[]User `gorm:"many2many:user_role;"`
}


//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&Role{})
//	//添加默认数据
//	initDBGorm.Db.Create(&Role{Id: 1,Name: "总后台超级管理员",PermissionType: 1})
//	initDBGorm.Db.Create(&Role{Id: 2,Name: "服务商超级管理员",PermissionType: 2})
//	initDBGorm.Db.Create(&Role{Id: 3,Name: "商户超级管理员",PermissionType: 3})
//}
//跟据用户id查询角色信息
func SelRolesByUserId(userId int64)(roles []Role,err error) {
	var rolesRes []Role
	err=initDBGorm.Db.Model(Role{}).
		Joins(" left join user_roles as b on b.role_id =roles.id").
		Where(" b.user_id=?",userId).Find(&rolesRes).Error
	if err!=nil{
		return rolesRes,err
	}
	return rolesRes,err
}

func (role *Role) AddRoleByGorm()int64 {
	//gin+gorm操作数据库方式
	data:=Role{
		Name: role.Name,
		Created_at:time.Now().Unix(),
		Updated_at:time.Now().Unix(),
	}
	result:=initDBGorm.Db.Create(&data)
	fmt.Println("输出")
	//fmt.Println(result.Error)//error
	fmt.Println(result.RowsAffected )//插入记录的条数
	fmt.Println(data.Id)//返回新创建的id
	return data.Id
}

func (role *Role)SelRoleByGorm()[]Role {
	var res[] Role
	name:=role.Name
	initDBGorm.Db.Where("name LIKE ?", "%"+name+"%").Find(&res)
	//template:="2006-01-02 15:04:05"
	//res2[index].Updated_at=time.Unix(int64(res[index].Updated_at),0).Format(template)
	return res
}
//转成ToRoleResponse
func (role Role)ToRoleResponse()RoleResponse  {

	return RoleResponse{
		Id  :role.Id,
		Name :role.Name,
		Created_at :utils.TimestampToTime(role.Created_at),
		Updated_at :utils.TimestampToTime(role.Updated_at),
	}
}