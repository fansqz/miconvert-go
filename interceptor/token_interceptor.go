package interceptor

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/result"
	"miconvert-go/utils"
)

//
// TokenAuthorize
//  @Description: token拦截器
//  @return gin.HandlerFunc
//
func TokenAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := result.NewResult(c)
		token := c.Request.Header.Get("token")
		user, err := utils.ParseToken(token)
		if err != nil || user == nil {
			r.Error(result.IDENTITY_INVALID.GetCode(),
				result.IDENTITY_INVALID.GetMessage(), nil)
		}
	}
}
