package main

import (
	"custom_controller/client"
	"custom_controller/controller"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	clientSet, err := client.GetClientSet()
	if err != nil {
		log.Fatalf("获取配置文件失败:%v", err)
	}

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	podListWatcher := cache.NewListWatchFromClient(clientSet.CoreV1().RESTClient(), "pods", "default", fields.Everything())
	indexer, informer := cache.NewIndexerInformer(podListWatcher, &corev1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			fmt.Printf("加入pod:%s\n", key) //拿到key
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			fmt.Printf("更新queue:%s\n", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			fmt.Printf("删除pod:%s\n", key)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	controller := controller.NewController(indexer, queue, informer)
	stopCh := make(chan struct{})
	controller.Run(1, stopCh)
	defer close(stopCh)
}
