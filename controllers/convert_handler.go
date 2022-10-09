package controllers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/dao"
	r "miconvert-go/result"
	"strings"
)

//
// ConvertController
// @Description: 用于做各种文件转换的handler
//
type ConvertController interface {
	//查询支持的类型转换
	GetSupportOutFormat(ctx *gin.Context)
	//添加文件进行解析
	ConvertFiles(ctx *gin.Context)
	//下载解析文件
	DownloadFiles(ctx *gin.Context)
}

type convertController struct {
}

func NewConvertController() *convertController {
	return &convertController{}
}

func (c *convertController) GetSupportOutFormat(ctx *gin.Context) {
	result := r.NewResult(ctx)
	inFileName := ctx.Param("fileNam")
	//检验文件名称是否合理
	a := strings.Split(inFileName, ".")
	if len(a) < 2 {
		result.SimpleErrorMessage("不支持该文件格式")
	}
	//根据格式获取支持转换类型
	outFormats, err := dao.ListAllOutFormatByInFormat(inFileName)
	if err != nil || len(outFormats) == 0 {
		result.SimpleErrorMessage("不支持该文件格式")
	}
	result.SuccessData(outFormats)
}

func (c *convertController) ConvertFiles(ctx *gin.Context) {

}

func (c *convertController) DownloadFiles(ctx *gin.Context) {

}
