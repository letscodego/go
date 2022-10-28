package main

import "fmt"

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["ManUnited"] = 20
	leagueTitles["Sunderland"] = 4
	leagueTitles["Newcastle"] = 6

	recentHead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle":  0,
	}

	fmt.Println(leagueTitles)
	fmt.Println(recentHead2HeadWins)

	fmt.Printf("Recent head 2 head wins is %v", recentHead2HeadWins)

}
