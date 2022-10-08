// Package utils
// @Author: fzw
// @Create: 2022/10/8
// @Description: 工具包
package utils

import (
	"log"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

//
// SOfficeConvert
//  @Description: 利用libreoffice进行格式转换
//  @param fileSrcPath 原文件位置
//  @param fileOutDir 文件保存位置
//  @param toFormat 转换的格式
//  @return err
//
func SOfficeConvert(fileSrcPath string, fileOutDir string, toFormat string) (fileOutPath string, err error) {
	//根据不同的系统，调用不同命令
	osname := runtime.GOOS
	command := ""
	if osname == "windows" {
		command = "soffice"
	}
	if osname == "linux" {
		command = "libreoffice6.0"
	}
	cmd := exec.Command(command, "--invisible", "--convert-to",
		toFormat, fileSrcPath, "--outdir", fileOutDir)
	byteByStat, errByCmdStart := cmd.Output()
	//错误则返沪
	if errByCmdStart != nil {
		err = errByCmdStart
		return
	}
	//成功准备数据返回
	filename := strings.Split(path.Base(fileSrcPath), ".")[0]
	fileOutPath = fileOutDir + "/" + filename + "." + toFormat
	log.Println("文件转换成功", string(byteByStat))
	return fileOutPath, nil
}

//
// Pdf2docxConvert
//  @Description: 利用pdf2docx进行转换，只能从pdf转word,pdf->docx
//  @param fileSrcPath 源文件路径
//  @param fileOutDir 导出文件保存位置
//  @return fileOutPath 导出文件路径
//  @return err
//
func Pdf2docxConvert(fileSrcPath string, fileOutDir string) (fileOutPath string, err error) {
	//获取输出文件名路径
	filename := strings.Split(path.Base(fileSrcPath), ".")[0]
	fileOutPath = fileOutDir + "/" + filename + ".docx"
	//执行
	cmd := exec.Command("pdf2docx", "convert", fileSrcPath, fileOutPath)
	byteByState, errByCmdStart := cmd.Output()
	if errByCmdStart != nil {
		err = errByCmdStart
		return
	}
	//返回数据
	log.Println("文件转换成功", string(byteByState))
	return fileOutPath, nil
}
