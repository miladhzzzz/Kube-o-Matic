package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	// vx "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conf "k8s.io/client-go/applyconfigurations/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeController struct {
	ctx context.Context
}

func NewKubeController() KubeController {
	return KubeController{ctx: context.TODO()}
}

func findKubeConfig(fs string) ([]string, error) {

	var files []string

	root := fs

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func kubectl() (*kubernetes.Clientset, error) {

	kubeConfigPath, err := findKubeConfig("/kubeconfig")

	if err != nil {
		return nil, err
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath[0])

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)

	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func (kc *KubeController) Deploy() gin.HandlerFunc {
	
    return func(ctx *gin.Context) {
        ns := ctx.Param("ns")

        // Get the YAML file from the request body
        var deployment conf.DeploymentApplyConfiguration

        if err := ctx.BindJSON(&deployment); err!= nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Create the Kubernetes client set
        clientSet, err := kubectl()
        if err!= nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		
        // Deploy the deployment to Kubernetes
        if _, err := clientSet.AppsV1().Deployments(ns).Apply(context.Background(), &deployment, v1.ApplyOptions{}); err!= nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{"message": "Deployment created successfully"})
    }
}

func (kc *KubeController) GetPods() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ns := ctx.Param("ns")

		clientset, err := kubectl()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, pods)

	}
	
}
