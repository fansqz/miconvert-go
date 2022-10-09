// Package result
// @Author: fzw
// @Create: 2022/10/8
// @Description: 用做统一结果返回的
package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// result
// @Description: 统一result，并返回数据给前端
//
type result struct {
	ctx *gin.Context
}

func NewResult(ctx *gin.Context) *result {
	return &result{ctx: ctx}
}

//
// Success1
//  @Description: 返回成功结果
//  @receiver r
//  @param data
//
func (r *result) Success1() {
	res := Success()
	r.ctx.JSON(http.StatusOK, res)
}

func (r *result) Success2(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Success()
	res.Data = data
	r.ctx.JSON(http.StatusOK, res)
}

func (r *result) Success3(message string) {
	res := &ResultCont{
		Code:    OK.GetCode(),
		Message: message,
	}
	r.ctx.JSON(http.StatusOK, res)
}

func (r *result) Success4(message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := &ResultCont{
		Code:    OK.GetCode(),
		Message: message,
		Data:    data,
	}
	r.ctx.JSON(http.StatusOK, res)
}

//
// Error1
//  @Description: 返回错误的消息
//  @receiver r
//
func (r *result) Error1() {
	res := Error()
	r.ctx.JSON(http.StatusOK, res)
}

func (r *result) Error2(code int, message string) {
	res := &ResultCont{
		Code:    code,
		Message: message,
	}
	r.ctx.JSON(http.StatusOK, res)
}

func (r *result) Error3(code int, message string, data interface{}) {
	res := &ResultCont{
		Code:    code,
		Message: message,
		Data:    data,
	}
	r.ctx.JSON(http.StatusOK, res)
}
