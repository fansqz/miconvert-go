package models

const (
	ACTIVATED   = 1
	INACTIVATED = 0
)

//User 结构体
type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	State    int    `gorm:"column:state"`
	Code     string `gorm:"column:code"`
}
