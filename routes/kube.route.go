package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miladhzzzz/kube-o-matic/controllers"
)

var (
	// ctx    *gin.Context
)

type KubeRouteController struct {
	kubeController controllers.KubeController
}

func NewKubeRouteController(kubeController controllers.KubeController) KubeRouteController {
	return KubeRouteController{kubeController}
}

func (rc *KubeRouteController) KubeRoute(rg *gin.RouterGroup) {

	router := rg.Group("/kubectl")


	rg.POST("/upload", func(c *gin.Context) {
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
	
	router.POST("/deploy/:ns", rc.kubeController.Deploy())

	router.GET("/pods/:ns", rc.kubeController.GetPods())

}