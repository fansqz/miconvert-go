package dao

import (
	"miconvert-go/db"
	"miconvert-go/models"
)

//GetUserByName 根据用户名读取一条数据
func GetUserByName(username string) (*models.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := db.DB.Raw(sqlStr, username)
	user := &models.User{}
	row.Scan(&user)
	return user, nil
}

// CheckUserNameInDb 检测用户名是否存在
func CheckUserNameInDb(username string) bool {
	sqlStr := "select name from users where username = ?"
	//执行
	row := db.DB.Raw(sqlStr, username)
	var username2 string
	row.Scan(&username2)
	return username2 != ""
}

// CheckEmailInDb 检测邮箱是否存在
func CheckEmailInDb(email string) bool {
	sqlStr := "select email from users where email = ?"
	row := db.DB.Raw(sqlStr, email)
	var email2 string
	row.Scan(&email2)
	return email2 != ""

}

//InsertUser 向数据库中插入用户信息
func InsertUser(user *models.User) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	db.DB.Exec(sqlStr, user.Username, user.Password, user.Email)
	return nil
}
