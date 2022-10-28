package main

import "fmt"

func main() {
	//courses := make([]string, 5, 10)

	courses := []string{"C#", "K8s", "Docker", "F#"}
	fmt.Printf("Len of slice is %d and capacity is %d\n", len(courses), cap(courses))

	courses = append(courses, "Go")
	for _, i := range courses {
		fmt.Println(i)
	}
}
