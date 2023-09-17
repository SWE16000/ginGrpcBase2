package service

import (
	"context"
	"encoding/json"
	"fmt"
	initDBGorm "ginGrpcBase2/database/gorm"
	"ginGrpcBase2/model"
	"ginGrpcBase2/pdgrpc"
	"ginGrpcBase2/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

//Grpc接口具体实现
type UserService struct {
}
//func NewUserService() *UserService {
//	return &UserService{}
//}


func (u  UserService)UserEdit(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	fmt.Println("通过grpc调用用户编辑接口")

	res:= pdgrpc.CommonResult{
		Code:200,
		Message:"asda",
	}
	return &res,nil
}
func (u  UserService)UserList(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.ListReponse, error){
	return nil,nil
}
func (u  UserService)TaskDetail(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	return nil,nil
}
func (u  UserService)UserDelete(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	return nil,nil
}
func (u  UserService)UserLogin(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	fmt.Println("通过grpc调用用户登录接口")
	var loginRequest pdgrpc.LoginRequest
	res:= &pdgrpc.CommonResult{
		Code:0,
		Message:"",
	}
	err := anypb.UnmarshalTo(request.Data, &loginRequest, proto.UnmarshalOptions{})
	if err!=nil{
		fmt.Println("CommonRequest")
	}
	//查询用户
	var users model.User
	err=initDBGorm.Db.Where("username=?",loginRequest.Username).First(&users).Error
	if err != nil {
		res.Code=400
		res.Message="当前用户不存在,请重新登录"
		return res,nil
	}
	//验证密码正确性
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(loginRequest.Password)); err != nil {
		res.Code=400
		res.Message="账号密码错误"
		return res,nil
	}
	var userResponse  model.UserResponse
	userResponse,err=model.SelectUserRoleById(users.Id)
	if err!=nil{
	}
	var pUserResponse pdgrpc.UserResponse
	pUserResponse.Id=userResponse.Id
	pUserResponse.Username=userResponse.Username
	pUserResponse.Phone=userResponse.Phone
	bytes, err := json.Marshal(userResponse.Roles)
	var roles []*pdgrpc.Role
	json.Unmarshal(bytes, &roles)
	for _,value := range roles {
		pUserResponse.Roles=append(pUserResponse.Roles,value)
	}
	//pUserResponse.Roles=append(pUserResponse.Roles,userResponse.Roles)
	marshal, err := anypb.New(&pUserResponse)
	if err!=nil{
		utils.Logger.Error("传递参数转成any类型失败："+err.Error())
		return nil,err
	}
	res.Code=200
	res.Message="登录成功"
	res.Data=marshal
	return res,nil
}
func (u  UserService)UserLout(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	return nil,nil
}
func (u UserService)UserRegister(ctx context.Context, request *pdgrpc.CommonRequest) (*pdgrpc.CommonResult, error){
	return nil,nil
}


