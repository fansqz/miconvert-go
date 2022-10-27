package models

//User 结构体
type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Username string `gorm:"colum:username"`
	Password string `gorm:"colum:password"`
	Email    string `gorm:"colum:email"`
}
