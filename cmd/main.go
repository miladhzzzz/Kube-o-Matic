package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/miladhzzzz/kube-o-matic/controllers"
	"github.com/miladhzzzz/kube-o-matic/routes"

	"github.com/gin-gonic/gin"
)

var (
	KubeController controllers.KubeController
	KubeRouteController routes.KubeRouteController
	serviceName = "kubeomatic"
	server *gin.Engine
)

func init () {

	KubeController = controllers.NewKubeController()

	KubeRouteController = routes.NewKubeRouteController(KubeController)

	logFile, _ := os.Create("kubeomatic-service-http.log")

	server = gin.Default()
	server.Use(gin.LoggerWithWriter(logFile))

}


func main() {
	// starting gin server
	startGinServer()
	
}

func startGinServer() {

	router := server.Group("")

	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8555"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// server.Use(otelgin.Middleware(serviceName))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "value"})
	})

	KubeRouteController.KubeRoute(router)

	log.Fatal(server.Run(":8555"))

}