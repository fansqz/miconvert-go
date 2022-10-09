// Package models
// @Author: fzw
// @Create: 2022/10/9
// @Description: 用于存放各种po，dto等
package models

import (
	"miconvert-go/utils"
)

//
// Format
// @Description: 记录格式，和格式转换
//
type Format struct {
	ID              string    `gorm:"primary_key;column:id"`                                                   //格式id
	Name            string    `gorm:"colum:name"`                                                              //格式类型
	LoSupportFormat []*Format `gorm:"many2many:lo_support_format;association_jointable_foreignkey:support_id"` //格式可以通过libreoffice转换为其他的格式
}

func init() {
	//使用时进行注册
	utils.DB.AutoMigrate(&Format{})
}
