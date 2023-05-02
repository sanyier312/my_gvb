package main

import (
	_ "my_gvb/docs"
	"my_gvb/models"
	"my_gvb/routers"
)

// @title   API文档
// @version 1.0
// @description   API文档描述
// @host   127.0.0.01:8080
// @BasePath /
func main() {
	models.InitDb() //引用数据库

	routers.InitRouter() //引入路由组件

}
