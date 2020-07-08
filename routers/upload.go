package routers

import (
	"github.com/gin-gonic/gin"
	v1controller "seifwu.com/gin-basic-project/controller/api/v1"
)

// UploadRoutes 路由
func UploadRoutes(r *gin.Engine) {
	r.POST("/api/v1/upload", v1controller.Upload)
}
