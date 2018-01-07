package controllers

import (
	"Blog-text/models"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	//"time"
)

func Text(c *gin.Context) {

	//c.String(http.StatusOK, "hello dc!")
	c.HTML(http.StatusOK, "text.html", gin.H{
		"msg": "ddd",
	})
}

func TextAdd(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")

	err := models.AddText(title, content)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "有错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"content": content,
	})

	//c.String(http.StatusOK, fmt.Sprintf(" title=%s and content=%s.", title, content))

}

func GetTexts(c *gin.Context) {
	//texts := make([]models.Text, 0)

	get, err := models.GetAllText()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询失败！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"texts": get,
	})
}

func GetText(c *gin.Context) {
	tid := c.PostForm("id")

	get, err := models.GetText(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"text": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"text": get,
	})

}

func TextUpdate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	tid := c.PostForm("id")

	err := models.ModifyText(tid, title, content)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功！",
	})

	get, err := models.GetText(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "更新失败",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "提取修改后文章成功！",
		"text": get,
	})
}

func DeleteText(c *gin.Context) {
	tid := c.PostForm("id")

	// if tid == "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "无法获取此数据，可能没有！",
	// 	})
	// 	return
	// }

	err1 := models.DeleteText(tid)
	if err1 != nil {
		log.Println(err1)
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除失败！",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功！",
	})

}
