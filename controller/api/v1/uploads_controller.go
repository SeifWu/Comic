package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/global/response"
	"seifwu.com/gin-basic-project/model"
	"seifwu.com/gin-basic-project/utils"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	var attachment model.Attachment
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		rootedPath, _ := os.Getwd()
		filename := utils.FileRename(file.Filename)
		if err := c.SaveUploadedFile(file, rootedPath+"/assets/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		qiniuResult := utils.QiNiuUpload(filename, rootedPath+"/assets/"+filename)

		if qiniuResult["hash"] != "exist" {
			attachment = model.Attachment{
				FileName: filename,
				FileSize: file.Size,
				SignID:   qiniuResult["hash"],
			}
			global.DB.Create(&attachment)
		}
	}

	response.Success(c, gin.H{"files": attachment}, "上传成功")
}
