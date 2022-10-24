package utils

import (
	"github.com/dgrijalva/jwt-go"
	"miconvert-go/models"
	"time"
)

const (
	//jwt的私钥
	key = "miconvert_key*.*"
	//过期时间
	expiredTime = 12 * time.Hour
)

//定义一个实体
type claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

//
// GenerateToken
// https://www.jianshu.com/p/202b04426368
//  @Description:  通过user生成token
//  @param user    用户
//  @return string 生成的token
//  @return error
//
func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(expiredTime)
	claims := claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expiredTime.Unix(),
			//指定token发行人
			Issuer: "gin-blog",
		},
	}
	//设置加密算法，生成token对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//通过私钥获取已签名token
	token, err := tokenClaims.SignedString([]byte(key))
	return token, err
}

//
// ParseToken
//  @Description: 解析token，返回user
//  @param token token
//  @return user 用户
//
func ParseToken(token string) (*models.User, error) {
	//获取到token对象
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	//通过断言获取到claim
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			user := &models.User{}
			user.ID = claims.ID
			user.Username = claims.Username
			user.Email = claims.Email
			return user, nil
		}
	}
	return nil, err
}
