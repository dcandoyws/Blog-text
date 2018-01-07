package routers

import (
	// "blogtest/a

	c "Blog-text/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//文章的创建读写删除
	router.GET("/home/txt", c.Text)
	router.POST("/home/txtcreate", c.TextAdd)
	router.GET("/home/txtreads", c.GetTexts)
	router.POST("/home/txtread", c.GetText)
	router.POST("/home/txtupdate", c.TextUpdate)
	router.POST("/home/txtdelete", c.DeleteText)
	//
	return router

}
