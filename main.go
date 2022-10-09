package main

import (
	"fmt"
	"miconvert-go/routers"
	"miconvert-go/setting"
	"miconvert-go/utils"
	"os"
)

func main() {
	//获取参数
	if len(os.Args) < 2 {
		fmt.Println("参数错误")
		return
	}
	path := os.Args[1]
	//加载配置
	if err := setting.Init(path); err != nil {
		fmt.Println("加载配置文件出错")
		return
	}
	//连接数据库
	if err := utils.InitMysql(setting.Conf.MySqlConfig); err != nil {
		fmt.Println("数据库连接失败")
	}
	//程序退出时关闭mysql
	defer utils.CloseMysql()
	//注册路由
	routers.Run()
}
