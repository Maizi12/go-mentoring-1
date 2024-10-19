package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

var (
	Routers = gin.Default()
)

func Setup() {
	Routers.Use(requestid.New())
	configs := cors.DefaultConfig()
	configs.AllowAllOrigins = true
	configs.AllowCredentials = true
	configs.AddAllowHeaders("token, api-key, Authorization, timestamps,xkey, user-agent,sec-ch-ua,sec-ch-ua-mobile,sec-ch-ua-platform,referer,golang-auth")
	Routers.Use(cors.New(configs))
	Routers.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Testing")

	})
}
