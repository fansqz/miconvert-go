// Package dao
// @Author: fzw
// @Create: 2022/10/9
// @Description: 用于存放各种对数据库操作的
package dao

import (
	"miconvert-go/db"
	"miconvert-go/models"
)

//
// InsertFormatConvert
//  @Description:  添加一个formatConvert
//  @param format
//  @return error
//
func InsertFormatConvert(formatConvert *models.FormatConvert) error {
	err := db.DB.Create(&formatConvert).Error
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
	if err = db.DB.Select("in_format").Find(&inFormats).Error; err != nil {
		return nil, err
	}
	return
}

//
// ListOutFormatByInFormat
//  @Description: 查询某个格式可以转换为其他格式
//  @param format
//  @return convertFormats
//  @return err
//
func ListOutFormatByInFormat(inFormat string) (outFormats []string, err error) {
	formatConverts := []*models.FormatConvert{}
	err = db.DB.Model(&models.FormatConvert{}).
		Select("out_format").Where("in_format = ?", inFormat).Scan(&formatConverts).Error
	if err != nil {
		return nil, err
	}
	if formatConverts != nil {
		outFormats = make([]string, len(formatConverts))
		for i := 0; i < len(formatConverts); i++ {
			outFormats[i] = formatConverts[i].OutFormat
		}
	}
	return
}

//
// ListAllOutFormat
//  @Description: 获取想要转换的格式
//  @param format
//  @return convertFormats
//  @return err
//
func ListAllOutFormat() (outFormats []string, err error) {
	formatConverts := []*models.FormatConvert{}
	err = db.DB.Model(&models.FormatConvert{}).
		Select("out_format").Scan(&formatConverts).Error
	if err != nil {
		return nil, err
	}
	if formatConverts != nil {
		outFormats = make([]string, len(formatConverts))
		for i := 0; i < len(formatConverts); i++ {
			outFormats[i] = formatConverts[i].OutFormat
		}
	}
	return
}

//
// ListInFormatByOufFormat
//  @Description: 根据输出格式，获取可转换的输入格式
//  @return inFormats
//  @return err
//
func ListInFormatByOufFormat(outFormat string) (inFormats []string, err error) {
	formatConverts := []*models.FormatConvert{}
	err = db.DB.Model(&models.FormatConvert{}).
		Select("in_format").Where("in_format = ?", outFormat).Scan(&formatConverts).Error
	if err != nil {
		return nil, err
	}
	if formatConverts != nil {
		inFormats = make([]string, len(formatConverts))
		for i := 0; i < len(formatConverts); i++ {
			inFormats[i] = formatConverts[i].InFormat
		}
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
	formatConvert := &models.FormatConvert{}
	err = db.DB.Where("in_format = ? and out_format = ?", intFormat, outFormat).Find(&formatConvert).Error
	if err != nil {
		return -1, err
	}
	if formatConvert == nil {
		return -1, nil
	}
	utilCode = formatConvert.ConvertUtil
	return
}
