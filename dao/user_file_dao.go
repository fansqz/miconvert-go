package dao

import (
	"miconvert-go/db"
	"miconvert-go/models"
)

//
// ListFileNamesByUserId
//  @Description: 通过用户名称获取用户所有文件名称
//  @param userId
//  @return []*models.UserFile
//
func ListFileNamesByUserId(userID int) []*models.UserFile {
	var userFiles []*models.UserFile
	db.DB.Select("id", "in_file_name", "out_file_name").Where("user_id = ?", userID).Scan(&userFiles)
	if userFiles == nil {
		userFiles = []*models.UserFile{}
	}
	return userFiles
}

//
// ListUserFileByUserId
//  @Description: 通过用户名称获取所有的userFile，包含所有属性
//  @param userId
//  @return []*models.UserFile
//
func ListUserFileByUserId(userID int) []*models.UserFile {
	var userFiles []*models.UserFile
	db.DB.Where("user_id = ?", userID).Scan(&userFiles)
	if userFiles == nil {
		userFiles = []*models.UserFile{}
	}
	return userFiles
}

//
// InsertUserFile
//  @Description: 添加一个用户文件
//  @param userFile
//
func InsertUserFile(userFile *models.UserFile) {
	db.DB.Create(userFile)
}

//
// DeleteUserFile
//  @Description: 删除一个用户文件
//  @param id
//
func DeleteUserFile(id int) {
	db.DB.Delete(&models.UserFile{}, 10)
}

//
// UpdateUserFile
//  @Description: 更新一个用户文件
//  @param userFile
//
func UpdateUserFile(userFile *models.UserFile) {
	db.DB.Model(&models.UserFile{}).Update(userFile)
}
