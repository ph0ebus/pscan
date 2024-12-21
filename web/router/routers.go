package router

import (
	"net/http"
	"pscan/web/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLFiles("./web/public/index.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/api/v1")
	{
		v1.GET("/index", controller.Index)
		v1.GET("/scan", controller.Scan)
	}
	return r
}
