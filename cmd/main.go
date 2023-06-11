package main

import (
	"os"
	"log"
	"net/http"

	"github.com/miladhzzzz/kube-o-matic/internal/k8s"

	"github.com/gin-gonic/gin"
)


func main() {
	// starting gin server
	startGinServer()
	
}

func startGinServer() {

	logFile, _ := os.Create("kubeomatic-service-http.log")

	server := gin.Default()

	router := server.Group("")

	router.Use(gin.LoggerWithWriter(logFile))

	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		log.Println(file.Filename)

		folder := "/kubeconfig/" + file.Filename

		err = os.MkdirAll(folder, 0755)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, folder)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"Upload was successfull": file.Filename})
	})

	router.GET("/pods/:ns", func(c *gin.Context) {
		
		pods, err := k8s.GetPods(c.Param("ns"))

		if err != nil {
			log.Printf("couldnt get pods: %v", err)
			c.JSON(http.StatusOK, err)
			return
		}

		c.JSON(http.StatusOK, pods.Items)

	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	log.Fatal(server.Run(":8555"))

}