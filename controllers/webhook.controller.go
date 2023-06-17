package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/miladhzzzz/kube-o-matic/internal/git"
	"github.com/gin-gonic/gin"
	webhook "github.com/go-playground/webhooks/v6/github"
)

type HookController struct {
	KubeController KubeController
	ctx context.Context
}

func NewHookcontroller(KubeController KubeController) HookController {
	return HookController{ctx: context.TODO(), KubeController: KubeController}
}

func(hc *HookController) HookHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Println("Received webhook...")

		webhookSecret := os.Getenv("WEBHOOK_SECRET")

		if webhookSecret == "" {
			log.Printf("please set your webhook secret via /webhook/secret/<YOUR SECRET HERE> endpoint. your webhook was recieved and was not proccesed.")
			return
		}

		hook, err := webhook.New(webhook.Options.Secret(webhookSecret))
		if err != nil {
			return
		}
		payload, e := hook.Parse(c.Request, webhook.PushEvent)
		if e != nil {
			log.Println("Error parsing", e)
		}

		switch payload.(type) {

		case webhook.PushPayload:

			event := payload.(webhook.PushPayload)
			url   := event.Repository.URL

			repo, err := git.NewGitRepository(url , false , "")

			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			
			dir := repo.Data()

			log.Println(dir)

			// hc.KubeController.Deploy()

		}
		c.Status(200)
	}
}

func(hc *HookController) SetSecret() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		secret := ctx.Param("secret")

		if secret == "" {
			ctx.JSON(http.StatusBadRequest, "you must provide your webhook secret")
			return
		}

		os.Setenv("WEBHOOK_SECRET", secret)

	}
}