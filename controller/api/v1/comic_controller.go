package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/global/response"
	"seifwu.com/gin-basic-project/model"
)

// ComicParams 创建漫画参数
type ComicParams struct {
	Title    string
	Describe string
	Author   string
	Cover    interface{}
}

// CreateComic 创建漫画
func CreateComic(c *gin.Context) {
	var comicParams ComicParams
	DB := global.DB

	if err := c.ShouldBindJSON(&comicParams); err != nil {
		log.Fatal(err)
	}

	newComic := model.Comic{
		Title: comicParams.Title,
	}

	DB.Create(&newComic)

	// 返回结果
	response.Success(c, nil, "创建成功")
}
