// Package setting
// @Author: fzw
// @Create: 2022/10/8
// @Description: 初始化时读取配置文件相关工具
package setting

import (
	"gopkg.in/ini.v1"
	"strings"
)

var Conf = new(AppConfig)

//
// AppConfig
// @Description:应用配置
//
type AppConfig struct {
	Release          bool   `ini:"release"` //是否是上线模式
	Port             int    `ini:"port"`    //端口
	ReleaseStartPath string `ini:"releaseStartPath"`
	*MySqlConfig
	*ConvertConfig
	*ReleasePathConfig
	*EmailConfig
}

//
// MySqlConfig
// @Description: mysql相关配置
//
type MySqlConfig struct {
	User     string `ini:"user"`     //用户名
	Password string `ini:"password"` //密码
	DB       string `ini:"db"`       //要操作的数据库
	Host     string `ini:"host"`     //host
	Port     string `ini:"port"`     //端口
}

//
// ConvertConfig
// @Description: 文件转换相关配置文件
//
type ConvertConfig struct {
	TempInPath  string `ini:"tempInPath"`  //输入文件临时存放位置
	TempOutPath string `ini:"tempOutPath"` //输出文件临时存放位置
	UserInPath  string `ini:"userInPath"`  //用户输入文件位置
	UserOutPath string `ini:"userOutPath"` //用户输出文件位置
}

type ReleasePathConfig struct {
	StartWith []string
}

type EmailConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

//
// Init
//  @Description: 初始化配置
//  @param file 配置文件路径
//  @return error
//
func Init(file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}
	mysqlConfig := &MySqlConfig{}
	convertConfig := &ConvertConfig{}
	emailConfig := &EmailConfig{}
	cfg.MapTo(Conf)
	cfg.Section("mysql").MapTo(mysqlConfig)
	cfg.Section("convert").MapTo(convertConfig)
	cfg.Section("email").MapTo(emailConfig)
	//遍历releasePath
	startPaths := strings.Split(Conf.ReleaseStartPath, ",")
	releasePathConfig := &ReleasePathConfig{StartWith: startPaths}
	Conf.ReleasePathConfig = releasePathConfig
	Conf.MySqlConfig = mysqlConfig
	Conf.ConvertConfig = convertConfig
	Conf.EmailConfig = emailConfig
	return nil
}
