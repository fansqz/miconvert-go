package interceptor

import (
	"github.com/gin-gonic/gin"
	result2 "miconvert-go/models/result"
	"miconvert-go/utils"
)

//
// TokenAuthorize
//  @Description: token拦截器
//  @return gin.HandlerFunc
//
func TokenAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := result2.NewResult(c)
		token := c.Request.Header.Get("token")
		user, err := utils.ParseToken(token)
		if err != nil || user == nil {
			r.Error(result2.IDENTITY_INVALID.GetCode(),
				result2.IDENTITY_INVALID.GetMessage(), nil)
		}
	}
}
