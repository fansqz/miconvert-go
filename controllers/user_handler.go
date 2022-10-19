package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"miconvert-go/dao"
	"miconvert-go/models"
	r "miconvert-go/result"
	"miconvert-go/utils"
)

//
// UserController
// @Description: 用户账号相关功能
//
type UserController interface {
	// Login 用户登录
	Login(ctx *gin.Context)
	// Logout 用户登出
	Logout(ctx *gin.Context)
	// Register 登出
	Register(ctx *gin.Context)
	// ChangePassword 改密码
	ChangePassword(ctx *gin.Context)
}

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) Register(ctx *gin.Context) {

}

func (u *userController) Login(ctx *gin.Context) {
	result := r.NewResult(ctx)
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	if username == "" {
		result.SimpleErrorMessage("用户名不可为空")
		return
	}
	if dao.CheckUserNameInDb(username) {
		result.SimpleErrorMessage("用户名称已存在")
		return
	}
	if len(password) < 6 {
		result.SimpleErrorMessage("密码不能小于6位")
	}
	//进行注册操作
	newPassword, err := utils.GetPwd(password)
	if err != nil {
		log.Println(err)
		result.SimpleErrorMessage("注册失败")
		return
	}
	user := &models.User{}
	user.Password = string(newPassword)
	user.Email = email
	user.Username = username
	//插入
	insertErr := dao.InsertUser(user)
	if insertErr != nil {
		log.Println(err)
		result.SimpleErrorMessage("注册失败")
		return
	}
	//注册成功返回数据
	result.SuccessMessage("注册成功")
}
