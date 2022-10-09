// Package models
// @Author: fzw
// @Create: 2022/10/9
// @Description: 用于存放各种po，dto等
package models

//
// FormatConvert
// @Description: 记录格式，和格式转换
//
type FormatConvert struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT;column:id"`        //格式id
	InFormat    string `gorm:"colum:in_format;unique_index:format_convert"` //格式类型
	OutFormat   string `gorm:"out_format;unique_index:format_convert"`      //可转换的格式类型
	ConvertUtil int    `gorm:"convert_util;unique_index:format_convert"`    //使用的工具
}
