package config

import (
	"fmt"
	"go-mentoring-1/app"
	"go-mentoring-1/pkg/tools"
	"go-mentoring-1/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	app.Setup(configPath)
	/*db.SetupDB()*/
	gin.SetMode(app.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = tools.CONFIGPATH
	}
	setConfiguration(configPath)
	conf := app.GetConfig()
	//sess := tools.ConnectAws(conf)
	/*redis := tools.InitializeRedis()*/
	//fmt.Println("SERVICE IS RUNNING ON " + conf.Server.Port)
	fmt.Printf("%+v\n", conf)
	router.Setup()
	gin.SetMode(conf.Server.Mode)
	defer RunAPI(":"+conf.Server.Port, router.Routers)

}

var (
	Conf = app.GetConfig()
)

func RunAPI(port string, router *gin.Engine) {

	servers := &http.Server{
		Handler: router,
	}
	fmt.Println("HTTP server is running on port " + port)

	servers.Addr = port
	if err := servers.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
