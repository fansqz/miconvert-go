package result

import "miconvert-go/constants"

//
// ResultCont
// @Description: unified response format
//
type ResultCont struct {
	Code    int         //response code
	Message string      //response message
	Data    interface{} //response data
}

//
// IsSuccessResult
//  @Description: 否是success
//  @param r 需要检验的resultCont
//  @return 是否success
//
func IsSuccessResult(r ResultCont) bool {
	return r.Code == constants.OK.GetCode()
}

//
// IsErrorResult
//  @Description: 是否是一个error的result
//  @param r 要检验的resultCont
//  @return 是否error
//
func IsErrorResult(r ResultCont) bool {
	return !IsSuccessResult(r)
}

//
// Success
//  @Description: 获取一个简单的成功响应的result
//  @return ResultCont
//
func Success() *ResultCont {
	return &ResultCont{
		Code:    constants.OK.GetCode(),
		Message: "",
	}
}

//
// Error
//  @Description: 获取一个简单的errorresult
//  @return ResultCont
//
func Error() *ResultCont {
	return &ResultCont{
		Code:    constants.Error.GetCode(),
		Message: constants.Error.GetMessage(),
	}
}
