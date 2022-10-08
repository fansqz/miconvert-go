package main

import (
	"fmt"
	"miconvert-go/db"
	"miconvert-go/models"
	"miconvert-go/routers"
	"miconvert-go/setting"
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
	if err := db.InitMysql(setting.Conf.MySqlConfig); err != nil {
		fmt.Println("数据库连接失败")
	}
	//程序退出时关闭mysql
	defer db.CloseMysql()
	//模型绑定,创建表
	db.DB.AutoMigrate(&models.Format{})
	//注册路由
	routers.Run()
}
