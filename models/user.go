package models

//User 结构体
type User struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username string `gorm:"colum:username" json:"username"`
	Password string `gorm:"colum:password" json:"password"`
	Email    string `gorm:"colum:email" json:"email"`
}
