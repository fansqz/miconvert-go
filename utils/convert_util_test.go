package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

//
// TestSOfficeConvert
//  @Description: 测试libreoffice转换工具
//  @param t
//
func TestSOfficeConvert(t *testing.T) {
	//获项目当前路径
	path, _ := os.Getwd()
	path = strings.ReplaceAll(path, "\\", "/")
	srcFilePath := path + "/resource/docxtest.docx"
	fileOutDir := path + "/resource"
	fileOutPath, err := SOfficeConvert(srcFilePath, fileOutDir, "pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fileOutPath)
}

//
// TestPdf2docxConvert
//  @Description: 测试pdf2docx转换工具
//  @param t
//
func TestPdf2docxConvert(t *testing.T) {
	//获项目当前路径
	path, _ := os.Getwd()
	path = strings.ReplaceAll(path, "\\", "/")
	srcFilePath := path + "/resource/pdftest.pdf"
	fileOutDir := path + "/resource"
	fileOutPath, err := Pdf2docxConvert(srcFilePath, fileOutDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fileOutPath)
}
