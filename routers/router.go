// Package routers
// @Author: fzw
// @Create: 2022/10/8
// @Description: 路由相关
package routers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/controllers"
	"miconvert-go/interceptor"
	"miconvert-go/setting"
	"miconvert-go/ws"
)

//
// Run
//  @Description: 启动路由
//
func Run() {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	//允许跨域
	r.Use(interceptor.Cors())
	//设置静态文件位置
	r.Static("/static", "/")
	//ping
	r.GET("/ping", controllers.Ping)
	//游客解析
	convert := r.Group("/convert")
	{
		convertController := controllers.NewConvertController()
		convert.POST("/convertFile", convertController.ConvertFile)
		convert.GET("/getSupportFormat", convertController.GetSupportOutFormat)
		convert.GET("/downloadFile/:filename", convertController.DownloadFile)
		convert.GET("/listAllOutFormat", convertController.ListAllOutFormat)
		convert.GET("/listAllInFormatByOutFormat", convertController.ListAllInFormatByOutFormat)
	}
	//用户注册,登录
	userController := controllers.NewUserController()
	r.POST("/user/register", userController.Register)
	r.POST("/user/login", userController.Login)

	//添加token拦截器
	r.Use(interceptor.TokenAuthorize())

	//修改密码
	r.POST("/user/changePassword", userController.ChangePassword)
	//ws
	r.GET("/ws/:token", func(ctx *gin.Context) {
		ws.ServeWs(ctx.Writer, ctx.Request)
	})
	//用户文件解析
	userConvert := r.Group("userConvert")
	{
		userConvertController := controllers.NewUserConvertController()
		userConvert.GET("listFile", userConvertController.ListFile)
		userConvert.DELETE("deleteFiles", userConvertController.DeleteFiles)
		userConvert.POST("convertFile", userConvertController.ConvertFile)
		userConvert.GET("downloadFile", userConvertController.DownloadFile)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
