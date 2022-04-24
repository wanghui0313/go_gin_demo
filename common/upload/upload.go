package upload

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const UPLOAD_DIR = "upload"

func Upload(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		}
		arr := strings.Split(file.Filename, ".")
		filePath, ok := checkExtAndgetFilePath(arr[1])
		if !ok {
			context.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  "ext invalid",
			})
			return
		}
		if err = context.SaveUploadedFile(file, filePath); err != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "upload success",
			"img_url": filePath,
		})
	}
}
func UploadView(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(200, "upload/index", gin.H{})
	}
}

func checkExtAndgetFilePath(fileExt string) (string, bool) {
	allowExt := map[string]bool{
		"jpg":  true,
		"png":  true,
		"jpeg": true,
	}
	if _, ok := allowExt[fileExt]; !ok {
		return "", false
	}
	if _, ok := os.Stat(UPLOAD_DIR); ok != nil {
		os.Mkdir(UPLOAD_DIR, os.ModePerm)
	}
	filePath := UPLOAD_DIR + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + fileExt
	return filePath, true
}

func UploadMul(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{"code": 400, "err": err.Error()})
			return
		}
		files := form.File["file"]
		var urls []string
		for _, file := range files {
			arr := strings.Split(file.Filename, ".")
			filePath, ok := checkExtAndgetFilePath(arr[1])
			if !ok {
				context.JSON(http.StatusForbidden, gin.H{
					"code": 400,
					"msg":  "ext invalid",
				})
				return
			}
			if err = context.SaveUploadedFile(file, filePath); err != nil {
				context.JSON(http.StatusForbidden, gin.H{
					"code": 400,
					"msg":  err.Error(),
				})
				return
			}
			urls = append(urls, filePath)
		}
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "upload  success", "img_url": urls})
	}
}
