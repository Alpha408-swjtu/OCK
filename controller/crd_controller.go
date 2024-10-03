package controller

import (
	"context"
	"custom_controller/client"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

var crd = schema.GroupVersionResource{
	Group:    "stable.example.com",
	Version:  "v1beta1",
	Resource: "redises",
}

func newListWatchForDynamicClient(dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, namespace string) *cache.ListWatch {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			obj, err := dynamicClient.Resource(crd).Namespace("default").List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return nil, err
			}
			return obj, nil
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			watchInterface, err := dynamicClient.Resource(gvr).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return nil, err
			}
			return watchInterface, nil
		},
	}
}

func NewCrdController() *CrdController {
	crdClient := client.NewCrdClient()
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	listWatcher := newListWatchForDynamicClient(crdClient, crd, "default")
	indexer, controller := cache.NewIndexerInformer(listWatcher, &unstructured.Unstructured{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			fmt.Printf("加入crd:%s\n", key) //拿到key
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			fmt.Printf("更新crd:%s\n", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			fmt.Printf("删除crd:%s\n", key)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})
	return &CrdController{
		indexer:  indexer,
		queue:    queue,
		informer: controller,
		client:   crdClient,
	}
}
