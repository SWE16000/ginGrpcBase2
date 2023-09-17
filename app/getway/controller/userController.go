package controller

import (
	"fmt"
	"ginGrpcBase2/app/getway/service"
	"ginGrpcBase2/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(context *gin.Context)  {
	fmt.Println("User Login")
	userRes, err:= service.UserLogin(context)
	if err != nil {
		utils.FailWithMsg(context,err.Error())
		return
	}
	utils.OkWithData(context,userRes)
	return
}