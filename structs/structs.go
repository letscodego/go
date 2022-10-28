package main

import "fmt"

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float32
	}

	//var gettingStartedWithK8s courseMeta
	//gettingStartedWithK8s := new(courseMeta)

	gettingStartedWithK8s := courseMeta{
		author: "Nigel Poulton",
		level:  "Intermediate",
		rating: 5,
	}

	fmt.Println("Author of the course is:", gettingStartedWithK8s.author)
	gettingStartedWithK8s.rating = 1.2
	fmt.Println("Rating of the course is:", gettingStartedWithK8s.rating)

}
