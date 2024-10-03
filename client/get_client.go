package client

import (
	"log"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func NewClientSet() *kubernetes.Clientset {
	clientSet, err := GetClientSet()
	if err != nil {
		log.Fatalf("获取配置文件失败:%v", err)
	}
	return clientSet
}

func NewCrdClient() *dynamic.DynamicClient {
	crdClient, err := GetCrdClient()
	if err != nil {
		log.Fatalf("获取配置文件失败:%v", err)
	}
	return crdClient
}
