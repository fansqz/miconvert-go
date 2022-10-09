package utils

import (
	"github.com/chilts/sid"
	"strings"
)

//
// ConvertToUniqueName
//  @Description: 把文件名转换为唯一文件名
//
func ConvertToUniqueName(filename string) string {
	name := sid.Id()
	a := strings.Split(filename, ".")
	format := a[1]
	return name + "." + format
}

func GetUUID() string {
	return sid.Id()
}
