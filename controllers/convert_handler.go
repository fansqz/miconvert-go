package controllers

import (
	"github.com/gin-gonic/gin"
)

//
// ConvertHandler
// @Description: 用于做各种文件转换的handler
//
type ConvertController interface {
	//查询支持的类型转换
	GetSupportFormat(ctx *gin.Context)
	//添加文件进行解析
	ConvertFiles(ctx *gin.Context)
	//下载解析文件
	DownloadFiles(ctx *gin.Context)
}

type convertController struct {
}

func NewConvertController() convertController {
	return convertController{}
}

func (c *convertController) GetSupportFormat(ctx *gin.Context) {

}

func (c *convertController) ConvertFiles(ctx *gin.Context) {

}

func (c *convertController) DownloadFiles(ctx *gin.Context) {

}
