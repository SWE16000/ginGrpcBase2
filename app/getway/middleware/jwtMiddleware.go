package middleware

//import (
//	"ginBase/database/redis"
//	"ginBase/model"
//	"ginBase/service"
//	"ginBase/utils"
//	"github.com/gin-gonic/gin"
//	"strings"
//)
//
/////JSON Web Token中间件
//func Jwt() gin.HandlerFunc{
//	return func(context *gin.Context){
//
//		//code := proto.SUCCESS
//		//获取请求头中的Authorization
//		authorization := context.Request.Header.Get("Authorization")
//		if authorization == "" {
//			utils.FailWithCode(context,401,"请求中未携带token,请重新登录")
//			context.Abort()
//			return
//		}
//		//拆分Authorization字段获取token字符串
//		strSlice := strings.SplitN(authorization, " ", 2)
//		if len(strSlice)!=2 && strSlice[0]!="Bearer"{
//			utils.FailWithCode(context,401,"token格式错误,请重新登录")
//			context.Abort()
//			return
//		}
//		//验证token字符串
//		claim,ok := service.ValidateJwt(strSlice[1])
//		if !ok {
//			utils.FailWithCode(context,401,"token格式错误,请重新登录")
//			context.Abort()
//			return
//		}
//		//过期判断
//		//if time.Now().Unix() > claim.ExpiresAt {
//		//	utils.FailWithCode(context,401,"token已过期,请重新登录")
//		//	context.Abort()
//		//	return
//		//}
//		flage, err :=redis.Db.Exists(context, claim.User.Username).Result()
//		if err!=nil{
//			utils.FailWithCode(context,401,"token失效,请重新登录")
//			context.Abort()
//			return
//		}
//		if flage==0{
//			utils.FailWithCode(context,401,"token已过期,请重新登录")
//			context.Abort()
//			return
//		}
//		//查看登录用户角色
//		var roleIds [] int64
//
//		roles,_:=model.SelRolesByUserId(claim.User.Id)
//		for _, value :=range roles{
//			roleIds=append(roleIds,value.Id)
//		}
//
//		//查看登录用户权限
//		//设置用户名
//		context.Set("username", claim.User.Username)
//		context.Set("token", strSlice[1])
//		context.Set("user_id", claim.User.Id)
//		context.Set("roles_id", roleIds)
//		context.Next()
//	}
//}