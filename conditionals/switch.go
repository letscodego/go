package main

import "fmt"

func main() {

	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1: Kubernetes with lower-case \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2: Kubernetes with a capital \"K\".")
	case "K8s":
		fmt.Println("Case 3: Kubernetes short name \"Kates\".")
	default:
		fmt.Println("Hit the default")
	}

}
