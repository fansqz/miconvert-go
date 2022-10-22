package controllers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/dao"
	r "miconvert-go/models/result"
	"miconvert-go/utils"
	"os"
	"strconv"
)

type UserConvertController interface {
	//获取用户文件列表
	ListFile(ctx *gin.Context)
	//用户批量删除文件
	DeleteFile(ctx *gin.Context)
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

func (c *userConvertController) DeleteFile(ctx *gin.Context) {
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
