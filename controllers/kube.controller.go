package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	vx "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

		clientSet, err := kubectl()

		if err != nil {
			log.Println(ns)
		}

		// clientSet.AppsV1().Deployments(ns).Apply(context.Background(), )

		log.Println(clientSet)

	}
}

func (kc *KubeController) GetPods(ns string) (*vx.PodList, error) {
	clientset, err := kubectl()

	if err != nil {
		return nil, err
	}

	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return pods, nil
}
