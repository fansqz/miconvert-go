// Package dao
// @Author: fzw
// @Create: 2022/10/9
// @Description: 用于存放各种对数据库操作的
package dao

import (
	"miconvert-go/models"
	"miconvert-go/utils"
)

//
// InsertFormatConvert
//  @Description:  添加一个formatConvert
//  @param format
//  @return error
//
func InsertFormatConvert(formatConvert *models.FormatConvert) error {
	err := utils.DB.Create(&formatConvert).Error
	return err
}

//
// ListAllInFormat
//  @Description: 获取输入文件可支持的所有格式
//  @param id
//  @return format
//  @return err
//
func ListAllInFormat() (inFormats []string, err error) {
	inFormats = []string{}
	if err = utils.DB.Select("in_format").Find(&inFormats).Error; err != nil {
		return nil, err
	}
	return
}

//
// ListAllOutFormatByInFormat
//  @Description: 查询某个格式可以转换为其他格式
//  @param format
//  @return convertFormats
//  @return err
//
func ListAllOutFormatByInFormat(inFormat string) (outFormats []string, err error) {
	outFormats = []string{}
	err = utils.DB.Select("out_format").Where("int_format = ?", inFormat).Find(&outFormats).Error
	if err != nil {
		return nil, err
	}
	return
}

//
// GetUtilByInFormatAndOutFormat
//  @Description: 通过输入格式和输出格式获取该转换可以使用的工具
//  @param intFormat
//  @param outFormat
//  @return utilCode
//  @return err
//
func GetUtilByInFormatAndOutFormat(intFormat string, outFormat string) (utilCode int, err error) {
	utilCode = -1
	err = utils.DB.Select("convert_util").Where("in_format = ? and out_format = ?", intFormat, outFormat).Find(&utilCode).Error
	if err != nil {
		return -1, err
	}
	return
}
