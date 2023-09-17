package model

import (
	initDBGorm "ginGrpcBase2/database/gorm"
	"ginGrpcBase2/pdgrpc"
	"ginGrpcBase2/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)
type User struct {
	Id   int64  `gorm:"primaryKey" form:"id"`
	Name string  `form:"name" gorm:"comment:用户姓名;size:255"`
	Username string `form:"username" gorm:"unique;comment:用户账号;size:255"`
	Password string  `form:"password" gorm:"comment:密码;size:255"`
	Phone string     `form:"phone" gorm:"comment:手机号;size:255"`
	IsDelete int64 `from:"is_delete" gorm:"comment:是否删除 1-未删除 2-已删除;default:1"`
	Created_at int64 `form:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	Updated_at int64 `form:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}



//func init()  {
//	//数据迁移创建数据表
//	initDBGorm.Db.AutoMigrate(&User{})
//	//添加默认数据
//	hash, _ :=bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
//	initDBGorm.Db.Create(&User{Id: 1,Name: "SWE16081",Username:"SWE16081",Phone: "15950643376",Password: string(hash)})
//}
//根据用户id查看用户角色信息
func SelectUserRoleById(id int64)(res UserResponse,err error)   {
	//获取用户信息
	var users User
	//构造用户角色数据
	var userResponse  UserResponse
	err=initDBGorm.Db.Model(User{}).Where("id=?",id).Find(&users).Error
	if err!=nil{
		return userResponse,err
	}
	var roles []map[string]interface{}
	err=initDBGorm.Db.Table("roles").Select("roles.*,user_roles.user_id").
		Joins("left join user_roles on roles.id=user_roles.role_id").Find(&roles).Error
	if err != nil {
		return userResponse,err
	}

		var resRoles [] Role
		for _,value2:=range roles{
			if users.Id==value2["user_id"]{
				resRoles=append(resRoles,Role{
					Id: value2["id"].(int64),
					Name:value2["name"].(string),
				})
			}
		}
		userResponse=users.ToUserResponse(resRoles)


	return userResponse,nil
}


func (user *User) AddUserByGorm()int64 {
	hash, err :=bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err.Error())
	}
	data:=User{
	Name: user.Name,
	Username: user.Username,
	Password: string(hash),
	Phone: user.Phone,
	Created_at:time.Now().Unix(),
	Updated_at:time.Now().Unix(),
	}
	result:=initDBGorm.Db.Create(&data)
	if result.Error!=nil{
		println("添加失败")
	}
	return data.Id
}



//转成UserResponse
func (user User)ToUserResponse(roles []Role)UserResponse  {

	return UserResponse{
		Id  :user.Id,
		Name :user.Name,
		Username :user.Username,
		Phone :user.Phone,
		Created_at :utils.TimestampToTime(user.Created_at),
		Updated_at :utils.TimestampToTime(user.Updated_at),
		Roles: roles,
	}
}
func (user User)ToUserResponse2(roles []pdgrpc.Role)pdgrpc.UserResponse  {
	res:=pdgrpc.UserResponse{
		Id  :user.Id,
		Name :user.Name,
		Username :user.Username,
		Phone :user.Phone,
		CreatedAt :utils.TimestampToTime(user.Created_at),
		UpdatedAt :utils.TimestampToTime(user.Updated_at),
	}
	for _,value := range roles {
		res.Roles=append(res.Roles,&value)
	}
	return res
}