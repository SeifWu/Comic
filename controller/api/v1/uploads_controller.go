package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"seifwu.com/gin-basic-project/global/response"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		rootedPath, _ := os.Getwd()
		filename := filepath.Base(file.Filename)
		log.Println("filename", filename)
		if err := c.SaveUploadedFile(file, rootedPath+"/assets/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	response.Success(c, gin.H{"files": files}, "上传成功")
}
