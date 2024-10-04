package controller

import (
	"custom_controller/client"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	Name     string
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller
	client   client.IClient
}

// 并发执行初始化函数
func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	fmt.Printf("启动%s-----------------\n", c.Name)
	go c.informer.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		fmt.Printf("等待缓存超时")
		return
	}
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.RunWoker, time.Second, stopCh)
	}
	<-stopCh
	fmt.Printf("关闭%s------------------\n", c.Name)
}

func (c *Controller) RunWoker() {
	for c.ProcessItem() {
	}
}

// 获取key
func (c *Controller) ProcessItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	if err := c.HandleObject(key.(string)); err != nil {
		if c.queue.NumRequeues(key) < 5 {
			c.queue.Add(key)
		}
	}
	return true
}

// 获取indexer里面的object
func (c *Controller) HandleObject(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Printf("事件出错:%s", key)
		return err
	}
	if !exists {
		fmt.Printf("Object :%s not found\n", key)
	} else {
		fmt.Printf(obj.(*corev1.Pod).GetName(), obj.(*corev1.Pod).GetNamespace())
		fmt.Println("成功！！！！！！！！")
	}

	return nil
}
