package controllers

import (
	"io"
	"log"
	"miconvert-go/dao"
	"miconvert-go/setting"
	"miconvert-go/utils"
	"mime/multipart"
	"os"
	"strings"
	"sync"
)

var ConvertManager *convertManager

const (
	CONVERTING = 0 //正在转换
	SUCCESS    = 1 //转换成功
	FALSE      = 2 //转换失败
)

//
// userFile
// @Description: 用于代表一个用户文件
//
type userFile struct {
	OriginalName string //文件原来的名称

	UniqueName string //数据库唯一名称

	UniqueConvertedName string //转换以后数据库唯一名称

	ConvertStat int //转换状态

	Message string //成功或者出错信息
}

//
// convertManager
// @Description: 用于管理和转换文件的
//
type convertManager struct {
	UserFileMap map[string][]*userFile //用户sessionID和该用户文件
	sync.RWMutex
}

//
// ReleaseSource
//  @Description: 用户断开ws连接，删除所有文件
//  @receiver c
//
func (c *convertManager) ReleaseSource(sessionId string) {
	//遍历删除用户文件
	userFiles := ConvertManager.UserFileMap[sessionId]
	c.Lock()
	for _, userFile := range userFiles {
		os.Remove(setting.Conf.TempInPath + "/" + userFile.UniqueName)
		if userFile.ConvertStat == SUCCESS {
			os.Remove(setting.Conf.TempOutPath + "/" + userFile.UniqueConvertedName)
		}
	}
	c.Unlock()
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
	c.Lock()
	userFiles := c.UserFileMap[sessionId]
	c.Unlock()
	return userFiles
}

//
// ConvertFile
//  @Description: 进行文件转换
//  @receiver c
//
func (c *convertManager) ConvertFile(sessionId string, file multipart.File, fileHead *multipart.FileHeader, outformat string) {
	outfilePath := ""
	userFile := &userFile{
		OriginalName: fileHead.Filename,
		ConvertStat:  CONVERTING,
	}
	defer func() {
		if outfilePath == "" {
			userFile.ConvertStat = FALSE
		}
		if _, err := os.Stat(outfilePath); err != nil {
			userFile.ConvertStat = FALSE
		}
		//通过路劲获取文件名
		outfilePath = strings.ReplaceAll(outfilePath, "\\", "/")
		a := strings.Split(outfilePath, "/")
		userFile.UniqueConvertedName = a[len(a)-1]
		userFile.ConvertStat = SUCCESS
	}()
	//存储到存储到临时文件
	c.Lock()
	c.UserFileMap[sessionId] = append(c.UserFileMap[sessionId], userFile)
	c.Unlock()
	defer file.Close()
	uniqueName := utils.ConvertToUniqueName(userFile.OriginalName)
	userFile.UniqueName = uniqueName
	os.Mkdir(setting.Conf.TempInPath, os.ModePerm)
	f, err := os.Create(setting.Conf.TempInPath + "/" + uniqueName)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//创建临时目录
	os.Mkdir(setting.Conf.TempOutPath, os.ModePerm)
	//获取可进行该格式转换的工具
	a := strings.Split(fileHead.Filename, ".")
	if len(a) != 2 {
		return
	}
	informat := a[1]
	utilCode, err := dao.GetUtilByInFormatAndOutFormat(informat, outformat)
	if err != nil {
		log.Println(err)
		return
	}
	if utilCode == utils.LIBRE_OFFICE {
		outfilePath, err = utils.SOfficeConvert(setting.Conf.TempInPath+"/"+userFile.UniqueName,
			setting.Conf.TempOutPath, outformat)
		if err != nil {
			return
		}
	}
	return
}

func init() {
	ConvertManager = &convertManager{
		UserFileMap: map[string][]*userFile{},
	}
}
