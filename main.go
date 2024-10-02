package main

import (
	"custom_controller/controller"
)

func main() {
	controller := controller.NewPodController()
	stopCh := make(chan struct{})
	controller.Run(1, stopCh)
	defer close(stopCh)
}
