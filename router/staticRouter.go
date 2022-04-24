package router

import "github.com/gin-gonic/gin"

func staticRouter(r *gin.Engine) {
	r.Static("img", "static/img")
	r.Static("js", "static/js")
	r.Static("css", "static/css")
	r.Static("upload", "upload")
}
