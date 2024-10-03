package controller

type Icontroller interface{
	Run(threadiness int, stopCh chan struct{})
	RunWoker()
	ProcessItem() bool
	HandleObject(key string)
}