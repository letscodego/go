package main

import "fmt"

func main() {

	courseInProg := []string{
		"Docker",
		"C#",
		"Golang",
		"K8s"}
	courseCompleted := []string{
		"C#"}

	for _, i := range courseInProg {
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println(i, "is completed.")
			}
		}
	}

}
