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
	//添加token拦截器
	r.Use(interceptor.TokenAuthorize())
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
		convert.GET("/listInFormatByOutFormat", convertController.ListInFormatByOutFormat)
		convert.GET("/listAllOutFormat", convertController.ListAllOutFormat)
	}
	//用户相关
	user := r.Group("/user")
	{
		userController := controllers.NewUserController()
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.POST("/user/changePassword", userController.ChangePassword)
	}
	//ws
	//r.GET("/ws/:token", func(ctx *gin.Context) {
	//	ws.ServeWs(ctx.Writer, ctx.Request)
	//})
	//用户文件解析
	userConvert := r.Group("/userConvert")
	{
		userConvertController := controllers.NewUserConvertController()
		userConvert.GET("/listFile", userConvertController.ListFile)
		userConvert.DELETE("/deleteFiles", userConvertController.DeleteFiles)
		userConvert.POST("/convertFile", userConvertController.ConvertFile)
		userConvert.GET("/downloadFile/:fileId", userConvertController.DownloadFile)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
