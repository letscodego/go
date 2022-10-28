package main

import "fmt"

func main() {
	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		fmt.Println("Docker Deep Dive is longer than caw course")
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Docker Deep Dive equals caw course")
	} else {
		fmt.Println("Docker Deep Dive is shorter than caw course")
	}
}
