package router

import (
	"github.com/gin-gonic/gin"
	"server/common/upload"
)

func uploadRouter(r *gin.Engine) {
	r.GET("/upload", upload.UploadView(r))
	r.POST("/upload", upload.Upload(r))
	r.POST("/uploadMul", upload.UploadMul(r))
}
