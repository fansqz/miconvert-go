package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"miconvert-go/dao"
	r "miconvert-go/result"
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
	//查询支持的类型转换
	GetSupportOutFormat(ctx *gin.Context)
	//添加文件进行解析,同步解析
	ConvertFile(ctx *gin.Context)
	//下载解析文件，同步解析时使用
	DownloadFile(ctx *gin.Context)
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
	outFormats, err := dao.ListAllOutFormatByInFormat(a[1])
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
		}
		if _, err := os.Stat(outfilePath); err != nil {
			result.SimpleErrorMessage("解析失败")
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
	//检验是否可以转换
	a := strings.Split(headler.Filename, ".")
	if len(a) != 2 {
		result.SimpleErrorMessage("不支持该文件格式")
	}
	inFormat := a[1]
	utilCode, err := dao.GetUtilByInFormatAndOutFormat(inFormat, outFormat)
	if err != nil {
		log.Println(err)
		result.SimpleErrorMessage("不支持该文件格式")
	}
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
	if utilCode == utils.LIBRE_OFFICE {
		outfilePath, err = utils.SOfficeConvert(infilePath,
			setting.Conf.TempOutPath, outFormat)
		if err != nil {
			log.Println(err)
			return
		}
	} else if utilCode == utils.PDF2DOCX {
		outfilePath, err = utils.Pdf2docxConvert(infilePath,
			setting.Conf.TempInPath)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (c *convertController) DownloadFile(ctx *gin.Context) {
	uniqueNameConverted := ctx.Query("key")
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
