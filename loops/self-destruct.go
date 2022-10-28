package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 10; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
