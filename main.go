package main

import (
	"bubble/model"
	"bubble/routers"

	"bubble/dao"
)

func main() {
	//创建数据库

	//链接数据库
	err := dao.InitMySQL()

	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库
	sqlDB, _ := dao.DB.DB()
	defer sqlDB.Close()

	//绑定模型，建立关联
	dao.DB.AutoMigrate(&model.Todo{})

	r := routers.SetUpRouter()
	r.Run()
}
