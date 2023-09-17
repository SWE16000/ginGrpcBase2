package service

import (
	"context"
	"errors"
	"fmt"
	"ginGrpcBase2/app/getway/reponse"
	"ginGrpcBase2/app/getway/request/userRequest"
	"ginGrpcBase2/model"
	"ginGrpcBase2/pdgrpc"
	"ginGrpcBase2/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

////用户登录
func UserLogin(ctx *gin.Context)(userRepose reponse.UserResponse,err error) {
	var request userRequest.LoginRequest
	userRepose=reponse.UserResponse{}
	var user model.User
	ctx.ShouldBind(&request)
	ctx.ShouldBind(&user)
	//验证参数
	msg := utils.Validate(&request, userRequest.LoginRequest{})
	if msg!=""{
		return userRepose, errors.New(msg)

	}
	prequest:=pdgrpc.LoginRequest{
		Username:request.Username,
		Password:request.Password,
	}
	//转换成any类型
	marshal, err := anypb.New(&prequest)
	if err!=nil{
		utils.Logger.Error("传递参数转成any类型失败："+err.Error())
		return userRepose,err
	}
	commonRequest :=&pdgrpc.CommonRequest{Data:marshal }
	// 从gin.Key中取出服务实例
	userService := ctx.Keys["user"].(pdgrpc.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), commonRequest)
	if err != nil {
		utils.Logger.Error("服务端出错，连接不上："+err.Error())
		return userRepose,err
	}
	var reponse pdgrpc.UserResponse
	err = anypb.UnmarshalTo(userResp.Data, &reponse, proto.UnmarshalOptions{})
	fmt.Println("userResp",reponse)

	if err!=nil{
		fmt.Println("CommonRequest")
	}
	//生成token
	if userResp.Code==200{
		token, err := GenToken(int(reponse.Id),ctx)
		if err!=nil{
			return userRepose,err
		}
		userRepose.User=reponse
		userRepose.Token=token
	}

	return userRepose,nil
}
