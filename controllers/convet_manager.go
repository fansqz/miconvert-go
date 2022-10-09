package controllers

import (
	"miconvert-go/setting"
	"os"
)

var ConvertManager *convertManager

//
// userFile
// @Description: 用于代表一个用户文件
//
type userFile struct {
	OriginalName string //文件原来的名称

	UniqueName string //数据库唯一名称

	UniqueConvertName string //转换以后数据库唯一名称

	ConvertSuccess bool //是否转换成功
}

//
// convertManager
// @Description: 用于管理用户文件的
//
type convertManager struct {
	UserFileMap map[string][]*userFile //用户sessionID和该用户文件

}

//
// ReleaseSource
//  @Description: 用户断开ws连接，删除所有文件
//  @receiver c
//
func (c *convertManager) ReleaseSource(sessionId string) {
	//遍历删除用户文件
	userFiles := ConvertManager.UserFileMap[sessionId]
	for _, userFile := range userFiles {
		os.Remove(setting.Conf.TempInPath + "/" + userFile.UniqueName)
		if userFile.ConvertSuccess {
			os.Remove(setting.Conf.TempOutPath + "/" + userFile.UniqueConvertName)
		}
	}
	//删除map
	delete(ConvertManager.UserFileMap, sessionId)
}

//
// ListConvertFile
//  @Description: 获取用户转换文件的信息
//  @receiver c
//  @param sessionId
//
func (c *convertManager) ListConvertFile(sessionId string) []*userFile {
	return c.UserFileMap[sessionId]
}

func init() {
	ConvertManager = &convertManager{
		UserFileMap: map[string][]*userFile{},
	}
}
