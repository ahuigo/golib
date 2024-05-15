package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func TestConfig(t *testing.T) {
	namespace := "dev"
	configname := "your-service"
	// kubeconfig: rancher donwload kubeconfig.yaml
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/k8s-kubeconfig.yaml")
	if err != nil {
		panic(err)
	}

	// 创建一个客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 读取一个 ConfigMap
	configmap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.Background(), configname, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data: %v\n", configmap.Data)

	// 修改 ConfigMap
	configmap.Data["new-key"] = "new-value"
	_, err = clientset.CoreV1().ConfigMaps("default").Update(context.Background(), configmap, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}
}
