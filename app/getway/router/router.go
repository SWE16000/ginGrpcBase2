package router

import (
	"ginGrpcBase2/app/getway/controller"
	"ginGrpcBase2/app/getway/middleware"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"strings"
)

func SetRouter(service ...interface{})*gin.Engine  {
	router:=gin.Default()
	router.Use(middleware.Cors(),middleware.InitMiddleware(service))
	router.Use(func(context *gin.Context) {
		rule := TestRule()
		context.String(http.StatusOK, rule)
		return
	})
	router.GET("/sel", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
		return
	})
	router.POST("/user/login", controller.UserLogin)
	//roleRouter := router.Group("/user",middleware.Jwt())
	//{
		//roleRouter.POST("/edit", controller.RoleEdit)
		//roleRouter.POST("/bindPermission", controller.RoleAddPermission)
		//roleRouter.GET("", controller.RoleIndex)
		//roleRouter.GET("/sel", controller.RoleSel)
		//roleRouter.DELETE("", controller.RoleDelete)
	//}


	return router
}


//测试流控规则
func TestRule()string {
	// 埋点（流控规则方式）
	_, b := sentinel.Entry("test", sentinel.WithTrafficType(base.Inbound))
	if b != nil {
		return "接口限流了！！！"
	}
	return "";
}
// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	totalFuncName := runtime.FuncForPC(pc[0]).Name()
	names := strings.Split(totalFuncName, ".")
	return names[len(names)-1]
}