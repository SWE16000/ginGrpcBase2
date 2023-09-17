package service

import (
	"errors"
	"fmt"
	"ginGrpcBase2/database/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)


type JWTClaims struct{
	jwt.StandardClaims
	UserId int
}
var (
	Secret     = "123#111" //salt
	ExpireTime = 3600      //token expire time
)
const (
	ErrorServerBusy = "server is busy"
	ErrorReLogin    = "relogin"
)

//生成 jwt token
//通过username生成token username要是唯一键
func GenToken(UserId int,context * gin.Context) (string, error) {
	claims := &JWTClaims{
		UserId: UserId,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	//token 存入redis
	_,err=redis.Db.Set(context,strconv.Itoa(UserId),signedToken, time.Second * time.Duration(ExpireTime)).Result()
	if err!=nil{
		return "",err
	}
	return signedToken, nil

}

//验证jwt token
func VerifyToken(tokenString  string) (*JWTClaims,error) {
	token, err := jwt.ParseWithClaims(tokenString , &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReLogin)
	}
	return claims, nil
}
//验证JSON Web Token
func ValidateJwt(tokenString string) (*JWTClaims, bool){
	//解析令牌字符串
	token,err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err!=nil {
		log.Println(err)
		return nil, false
	}
	//获取载荷
	claims,ok := token.Claims.(*JWTClaims)
	if ok && token.Valid{
		return claims, true
	}
	return nil, false
}
//token失效
func TokenDel(context *gin.Context,username string)error{
	_,err:=redis.Db.Del(context,username).Result()
	if err!=nil{
		fmt.Println("登出失败")
		fmt.Println(err)
		return errors.New("登出失败")
	}
	return nil
}


//// 更新token
//func Refresh(c *gin.Context) ( string,error) {
//	claims, _ := VerifyToken(c)
//	return GenToken(claims.User)
//}
//func JwtAuth(ctx *gin.Context) {
//	if _, err := VerifyToken(ctx); err == nil {
//		ctx.Next()
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{"code": 4001})
//		ctx.Abort()
//	}
//}
