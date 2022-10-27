package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"miconvert-go/dao"
	"miconvert-go/models"
	"miconvert-go/models/dto"
	r "miconvert-go/models/result"
	"miconvert-go/setting"
	"miconvert-go/utils"
	"miconvert-go/ws"
	"os"
	"strconv"
	"strings"
	"time"
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
	user := ctx.Keys["user"].(*models.User)
	userFiles := dao.ListFileStatesByUserId(user.Id)
	fileDtos := make([]*dto.FileDto, len(userFiles))
	//遍历并转换为fileDto
	for i := 0; i < len(userFiles); i++ {
		fileDto := &dto.FileDto{}
		fileDto.Id = userFiles[i].Id
		fileDto.UserId = userFiles[i].UserId
		fileDto.State = userFiles[i].State
		if userFiles[i].State == 2 {
			//转换成功
			fileDto.FileName = userFiles[i].OutFileName
			fileDto.FileSize = userFiles[i].OutFileSize
		} else {
			fileDto.FileName = userFiles[i].InFileName
			fileDto.FileSize = userFiles[i].InFileSize
		}
		fileDtos[i] = fileDto
	}
	result.SuccessData(fileDtos)
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
		dao.DeleteUserFile(userFile.Id)
	}
	result.SuccessMessage("删除成功!")
}

func (c *userConvertController) ConvertFile(ctx *gin.Context) {
	result := r.NewResult(ctx)
	token := ctx.GetHeader("token")
	user, _ := utils.ParseToken(token)
	outFormat := ctx.PostForm("outFormat")
	userFile := &models.UserFile{UserId: user.Id}
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
	userFile.InFileSize = utils.FormatFileSize(head.Size)
	userFile.State = models.CONVERTING
	userFile.Date = time.Now()
	dao.InsertUserFile(userFile)
	//进行异步解析
	go func() {
		outFilePath, convertErr := utils.Convert(infilePath, setting.Conf.UserOutPath, outFormat)
		if convertErr != nil {
			userFile.State = models.FALSE
		} else {
			userFile.OutFilePath = outFilePath
			//读取文件名称
			a := strings.Split(outFilePath, "/")
			userFile.OutFileName = a[len(a)-1][strings.Index(a[len(a)-1], "_")+1:]
			//读取输出文件大小
			osStat, statErr := os.Stat(outFilePath)
			if statErr != nil {
				log.Println(statErr)
			}
			userFile.OutFileSize = utils.FormatFileSize(osStat.Size())
			userFile.State = models.SUCCESS
		}
		//通过ws发送信息给用户
		dao.UpdateUserFile(userFile)
		//发送数据给前端
		userFiles := dao.ListUserFileByUserId(user.Id)
		json, _ := json.Marshal(r.ResultCont{
			Code: 200,
			Data: userFiles,
		})
		ws.WSManager.SendMessage(user.Id, json)
	}()
	result.SuccessData("文件已添加")
}

func (c *userConvertController) DownloadFile(ctx *gin.Context) {
	result := r.NewResult(ctx)
	fileIdString := ctx.Param("fileId")
	fileId, err := strconv.Atoi(fileIdString)
	if err != nil {
		result.SimpleErrorMessage("输入数据有误")
		return
	}
	//读取输出文件位置
	userFile := dao.GetUserFileById(fileId)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+userFile.OutFileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(userFile.OutFilePath)
	return
}
