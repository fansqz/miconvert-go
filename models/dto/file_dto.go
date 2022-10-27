package dto

import "time"

// FileDto 用于前端显示的文件，简化inpuFileName等，只有一个文件名一个文件大小
type FileDto struct {
	Id       int       `json:"id"`
	UserId   int       `json:"userId"`
	FileName string    `json:"fileName"`
	FilePath string    `json:"filePath"`
	FileSize string    `json:"fileSize"`
	State    int       `json:"state"`
	Date     time.Time `json:"date"`
}
