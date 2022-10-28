package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch tempNum := random(); tempNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got the following even number -", tempNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got the following odd number -", tempNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
