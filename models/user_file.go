package models

import "time"

const (
	CONVERTING = 0 //正在转换
	SUCCESS    = 1 //转换成功
	FALSE      = 2 //转换失败
)

//UserFile 结构体
type UserFile struct {
	Id          int       `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId      int       `gorm:"column:user_id"`
	InFileName  string    `gorm:"column:in_file_name"`
	OutFileName string    `gorm:"column:out_file_name"`
	InFilePath  string    `gorm:"column:in_file_path"`
	OutFilePath string    `gorm:"column:out_file_path"`
	InFileSize  string    `gorm:"column:in_file_size"`
	OutFileSize string    `gorm:"column:out_file_size"`
	State       int       `gorm:"column:state"`
	Date        time.Time `gorm:"column:date"`
}
