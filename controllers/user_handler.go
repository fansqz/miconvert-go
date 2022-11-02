package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"miconvert-go/dao"
	"miconvert-go/models"
	r "miconvert-go/models/result"
	"miconvert-go/utils"
)

//
// UserController
// @Description: 用户账号相关功能
//
type UserController interface {
	// Login 用户登录
	Login(ctx *gin.Context)
	// Register 注册
	Register(ctx *gin.Context)
	// 根据token获取用户信息
	GetUserInfo(ctx *gin.Context)
	// ChangePassword 改密码
	ChangePassword(ctx *gin.Context)
}

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) Register(ctx *gin.Context) {
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
	user.State = 0
	user.Code = utils.GetUUID()[0:6]
	//插入
	dao.InsertUser(user)
	//发送邮箱
	err = u.sendActivateEmail(user.Email, user.Code)
	//注册成功返回数据
	if err != nil {
		result.SimpleErrorMessage("注册失败，未知错误")
	} else {
		result.SuccessMessage("注册成功，激活链接将发送至邮箱，请点击激活")
	}
}

func (u *userController) Login(ctx *gin.Context) {
	result := r.NewResult(ctx)
	//获取并检验用户参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" {
		result.SimpleErrorMessage("用户名不可为空")
		return
	}
	if password == "" {
		result.SimpleErrorMessage("密码不可为空")
		return
	}
	user, userErr := dao.GetUserByName(username)
	if user.State == models.INACTIVATED {
		result.SimpleErrorMessage("用户未激活，请及时激活")
		return
	}
	if userErr != nil {
		result.SimpleErrorMessage("系统错误")
		log.Println(userErr)
		return
	}
	if user == nil || !utils.ComparePwd(user.Password, password) {
		result.SimpleErrorMessage("用户名或密码错误")
		return
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		log.Println(err)
		result.SimpleErrorMessage("系统错误")
		return
	}
	result.SuccessData(token)

}

func (u *userController) Activate(ctx *gin.Context) {
	result := r.NewResult(ctx)
	code := ctx.Param("code")
	err := dao.Activate(code)
	if err != nil {
		result.SimpleErrorMessage("激活失败")
	} else {
		result.SuccessMessage("激活成功")
	}
}

func (u *userController) ChangePassword(ctx *gin.Context) {
	result := r.NewResult(ctx)
	username := ctx.PostForm("username")
	oldPassword := ctx.PostForm("oldPassword")
	newPassword := ctx.PostForm("newPassword")
	if username == "" {
		result.SimpleErrorMessage("用户名不可为空")
		return
	}
	if oldPassword == "" {
		result.SimpleErrorMessage("请输入原始密码")
		return
	}
	//检验用户名
	user, err := dao.GetUserByName(username)
	if err != nil {
		log.Println(err)
		result.SimpleErrorMessage("系统错误")
		return
	}
	if user == nil {
		result.SimpleErrorMessage("用户名不存在")
		return
	}
	//检验旧密码
	if !utils.ComparePwd(oldPassword, user.Password) {
		result.SimpleErrorMessage("原始密码输入错误")
		return
	}
	password, getPwdErr := utils.GetPwd(newPassword)
	if getPwdErr != nil {
		result.SimpleErrorMessage("系统错误")
		log.Println(getPwdErr)
		return
	}
	user.Password = string(password)
	_ = dao.UpdateUser(user)
	token, daoErr := utils.GenerateToken(user)
	if daoErr != nil {
		result.SimpleErrorMessage("登录失败")
		return
	}
	result.SuccessData(token)
}

func (u *userController) GetUserInfo(ctx *gin.Context) {
	result := r.NewResult(ctx)
	user := ctx.Keys["user"].(*models.User)
	result.SuccessData(user)
}

func (u *userController) sendActivateEmail(email string, code string) error {
	message := "这是一封激活邮箱，点击链接激活miconvnert账号，如果不是你本人注册，请忽视该本条邮件\n " +
		"http://localhost:8080/user/activate/" + code
	return utils.SendMail([]string{email}, "miconvert激活邮箱", message)
}
