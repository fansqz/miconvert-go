package controllers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/dao"
	r "miconvert-go/models/result"
	"miconvert-go/utils"
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
