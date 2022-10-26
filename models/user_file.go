package models

import "time"

const (
	CONVERTING = 0 //正在转换
	SUCCESS    = 1 //转换成功
	FALSE      = 2 //转换失败
)

//User 结构体
type UserFile struct {
	ID          int       `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserID      int       `gorm:"colum:user_id"`
	InFileName  string    `gorm:"colum:in_file_name"`
	OutFileName string    `gorm:"colum:out_file_name"`
	InFilePath  string    `gorm:"colum:in_file_path"`
	OutFilePath string    `gorm:"colum:out_file_path"`
	InFileSize  string    `gorm:"colum:in_file_size"`
	OutFileSize string    `gorm:"colum:out_file_size"`
	State       int       `gorm:"colum:state"`
	Date        time.Time `gorm:"colum:date;date"`
}
