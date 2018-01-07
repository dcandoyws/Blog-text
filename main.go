package main

import (
	"Blog-text/models"
	"Blog-text/routers"
	"github.com/astaxie/beego/orm"
	//"github.com/gin-gonic/gin"
)

func Init() {
	models.InitDb()
}

func main() {

	Init()
	orm.Debug = true
	orm.RunSyncdb("default", false, false)

	//路由部分
	router := routers.InitRouter()
	//静态资源
	router.Static("/Static", "./static")
	//运行的端口
	router.Run(":8080")

}
