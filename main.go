package main

import "custom_controller/controller"

func main() {
	// controller1 := controller.NewCrdController()
	// stopch1 := make(chan struct{})
	// controller1.Run(1, stopch1)
	// defer close(stopch1)
	podContro := controller.NewPodController()
	stopch := make(chan struct{})
	podContro.Run(1, stopch)
	defer close(stopch)
}
