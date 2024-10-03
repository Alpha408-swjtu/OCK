package client

import (
	"log"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var configPath = "D:/kubeconfig/config"

func GetClientSet() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v", err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	return clientSet, err
}

func GetCrdClient() (*dynamic.DynamicClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v", err)
	}
	crdClient, err := dynamic.NewForConfig(config)
	return crdClient, err
}
