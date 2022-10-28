package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("Username")
	course := "Getting started with Kubernetes"

	fmt.Println("\nHi", name, "your current course is", course)
	UpdateCourse(&course)

	fmt.Println("You're currently watching", course)
}

func UpdateCourse(course *string) string {
	*course = "Getting Started with Docker"
	fmt.Println("Update course to", *course)
	return *course
}
