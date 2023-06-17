package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/miladhzzzz/kube-o-matic/controllers"
)

type HookRouteController struct {
// 	authController   controllers.AuthController
	hookController controllers.HookController
}

func NewHookRouteController(HookController controllers.HookController) HookRouteController {
	return HookRouteController{hookController: HookController}
}

func (rc *HookRouteController) HookRoute(rg *gin.RouterGroup) {
	
	router := rg.Group("")

	router.POST("/webhook", rc.hookController.HookHandler())

}
