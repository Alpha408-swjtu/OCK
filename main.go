package main

import "custom_controller/controller"

func main() {
	// controller = controller.NewCrdController()
	// stopch := make(chan struct{})
	// controller.Run(1, stopch)
	// defer close(stopch)
	cContro := controller.NewCrdController()
	stopch := make(chan struct{})
	cContro.Run(1, stopch)
	defer close(stopch)
}
