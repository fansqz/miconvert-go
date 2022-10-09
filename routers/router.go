// Package routers
// @Author: fzw
// @Create: 2022/10/8
// @Description: 路由相关
package routers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/controllers"
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
	//设置静态文件位置
	r.Static("/static", "/")

	r.GET("/ping", controllers.Ping)
	//文件转换相关
	convert := r.Group("/convert")
	{
		convertController := controllers.NewConvertController()
		convert.POST("/convertFiles", convertController.ConvertFiles)
		convert.GET("/getSupportFormat", convertController.GetSupportFormat)
		convert.GET("/downloadFiles", convertController.DownloadFiles)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
