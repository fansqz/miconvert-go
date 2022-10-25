package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"miconvert-go/dao"
	r "miconvert-go/models/result"
	"miconvert-go/setting"
	"miconvert-go/utils"
	"os"
	"strings"
	"time"
)

//
// ConvertController
// @Description: 用于做各种文件转换的handler
//
type ConvertController interface {
	// GetSupportOutFormat 查询支持的类型转换
	GetSupportOutFormat(ctx *gin.Context)
	// ConvertFile 添加文件进行解析,同步解析
	ConvertFile(ctx *gin.Context)
	// DownloadFile 下载解析文件，同步解析时使用
	DownloadFile(ctx *gin.Context)
	// ListAllOutFormat 获取所有可转换的格式
	ListAllOutFormat(ctx *gin.Context)
	// ListAllInFormatByOutFormat 根据输出格式，获取所有支持的输入格式
	ListInFormatByOutFormat(ctx *gin.Context)
}

type convertController struct {
}

func NewConvertController() *convertController {
	return &convertController{}
}

func (c *convertController) GetSupportOutFormat(ctx *gin.Context) {
	result := r.NewResult(ctx)
	inFileName := ctx.Query("fileName")
	//检验文件名称是否合理
	a := strings.Split(inFileName, ".")
	if len(a) < 2 {
		result.SimpleErrorMessage("不支持该文件格式")
		return
	}
	//根据格式获取支持转换类型
	outFormats, err := dao.ListOutFormatByInFormat(a[1])
	if err != nil || len(outFormats) == 0 {
		result.SimpleErrorMessage("不支持该文件格式")
		return
	}
	result.SuccessData(outFormats)
}

func (c *convertController) ConvertFile(ctx *gin.Context) {
	result := r.NewResult(ctx)
	outfilePath := ""
	infilePath := ""
	//最后判断解析是否成功
	defer func() {
		if outfilePath == "" {
			result.SimpleErrorMessage("解析失败")
			c.deleteSource(1*time.Second, infilePath)
			return
		} else if _, err := os.Stat(outfilePath); err != nil {
			result.SimpleErrorMessage("解析失败")
			c.deleteSource(1*time.Second, infilePath)
			return
		}
		//通过路劲获取文件名
		outfilePath = strings.ReplaceAll(outfilePath, "\\", "/")
		//延时10分钟删除资源
		c.deleteSource(10*time.Minute, infilePath)
		c.deleteSource(10*time.Minute, outfilePath)
		a := strings.Split(outfilePath, "/")
		result.SuccessData(a[len(a)-1])
	}()
	//获取文件名称
	file, headler, err := ctx.Request.FormFile("file")
	defer file.Close()
	outFormat := ctx.PostForm("outFormat")
	//保存文件到temp
	infilename := utils.GetUUID() + "_" + headler.Filename
	infilePath = setting.Conf.TempInPath + "/" + infilename
	os.Mkdir(setting.Conf.TempInPath, os.ModePerm)
	f, err := os.Create(infilePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//进行转换
	outfilePath, _ = utils.Convert(infilePath, setting.Conf.TempOutPath, outFormat)
}

func (c *convertController) DownloadFile(ctx *gin.Context) {
	uniqueNameConverted := ctx.Param("filename")
	ctx.Header("Content-Type", "application/octet-stream")
	filename := string(uniqueNameConverted[strings.Index(uniqueNameConverted, "_")+1:])
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(setting.Conf.TempOutPath + "/" + uniqueNameConverted)
	return
}

func (c *convertController) deleteSource(d time.Duration, path string) {
	go func() {
		t := time.NewTimer(d)
		<-t.C
		//删除资源
		os.Remove(path)
	}()
}

func (c *convertController) ListAllOutFormat(ctx *gin.Context) {
	result := r.NewResult(ctx)
	outFormats, err := dao.ListAllOutFormat()
	if err != nil || len(outFormats) == 0 {
		result.SimpleErrorMessage("系统错误")
		return
	}
	result.SuccessData(outFormats)
}

func (c *convertController) ListInFormatByOutFormat(ctx *gin.Context) {
	result := r.NewResult(ctx)
	outFormat := ctx.Query("outFormat")
	//根据格式获取支持转换类型
	inFormats, err := dao.ListInFormatByOufFormat(outFormat)
	if err != nil || len(inFormats) == 0 {
		result.SimpleErrorMessage("不支持该文件格式")
		return
	}
	result.SuccessData(inFormats)
}
