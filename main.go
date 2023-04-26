package main

import (
	"my_gvb/models"
	"my_gvb/routers"
)

func main() {
	models.InitDb() //引用数据库

	routers.InitRouter() //引入路由组件

}
