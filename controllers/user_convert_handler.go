package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"miconvert-go/dao"
	"miconvert-go/models"
	r "miconvert-go/models/result"
	"miconvert-go/setting"
	"miconvert-go/utils"
	"miconvert-go/ws"
	"os"
	"strconv"
	"strings"
)

type UserConvertController interface {
	//获取用户文件列表
	ListFile(ctx *gin.Context)
	//用户批量删除文件
	DeleteFiles(ctx *gin.Context)
	//ConvertFile 添加文件进行解析,同步解析
	ConvertFile(ctx *gin.Context)
	// DownloadFile 下载解析文件，同步解析时使用
	DownloadFile(ctx *gin.Context)
}

type userConvertController struct {
	convertController ConvertController
}

func NewUserConvertController() *userConvertController {
	return &userConvertController{}
}

func (c *userConvertController) ListFile(ctx *gin.Context) {
	result := r.NewResult(ctx)
	token := ctx.GetHeader("token")
	user, err := utils.ParseToken(token)
	if err != nil {
		result.SimpleErrorMessage("系统错误")
		return
	}
	userFiles := dao.ListFileNamesByUserId(user.ID)
	result.SuccessData(userFiles)
}

func (c *userConvertController) DeleteFiles(ctx *gin.Context) {
	result := r.NewResult(ctx)
	stringIDs := ctx.QueryArray("ids")
	//遍历每一个ids删除文件
	ids := make([]int, len(stringIDs))
	index := 0
	for i := 0; i < len(stringIDs); i++ {
		id, err := strconv.Atoi(stringIDs[i])
		if err == nil {
			ids[index] = id
			index++
		}
	}
	userFiles := dao.ListUserFileByIds(ids)
	//遍历文件进行删除
	for _, userFile := range userFiles {
		os.Remove(userFile.InFilePath)
		os.Remove(userFile.OutFilePath)
		dao.DeleteUserFile(userFile.ID)
	}
	result.SuccessMessage("删除成功!")
}

func (c *userConvertController) ConvertFile(ctx *gin.Context) {
	result := r.NewResult(ctx)
	token := ctx.GetHeader("token")
	user, _ := utils.ParseToken(token)
	outFormat := ctx.PostForm("outFormat")
	userFile := &models.UserFile{UserID: user.ID}
	//保存文件到temp
	file, head, err := ctx.Request.FormFile("file")
	defer file.Close()
	infilename := utils.GetUUID() + "_" + head.Filename
	infilePath := setting.Conf.UserInPath + "/" + infilename
	os.Mkdir(setting.Conf.UserInPath, os.ModePerm)
	f, err := os.Create(infilePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//添加到dao
	userFile.InFileName = head.Filename
	userFile.InFilePath = infilePath
	userFile.OutFileName = strings.Split(head.Filename, ".")[0] + "." + outFormat
	userFile.State = models.CONVERTING
	dao.InsertUserFile(userFile)
	//进行异步解析
	go func() {
		outFilePath, convertErr := utils.Convert(infilePath, setting.Conf.UserOutPath, outFormat)
		if convertErr != nil {
			userFile.OutFilePath = outFilePath
			userFile.State = models.FALSE
		} else {
			userFile.OutFileName = outFilePath
			userFile.State = models.SUCCESS
		}
		//通过ws发送信息给用户
		dao.UpdateUserFile(userFile)
		//发送数据给前端
		userFiles := dao.ListUserFileByUserId(user.ID)
		json, _ := json.Marshal(r.ResultCont{
			Code: 200,
			Data: userFiles,
		})
		ws.WSManager.SendMessage(user.ID, json)
	}()
	result.SuccessData("文件已添加")
}
