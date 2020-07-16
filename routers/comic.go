package routers

import (
	"github.com/gin-gonic/gin"
	v1controller "seifwu.com/gin-basic-project/controller/api/v1"
)

// ComicRoutes 路由
func ComicRoutes(r *gin.Engine) {
	r.POST("/api/v1/comic", v1controller.CreateComic)
}
