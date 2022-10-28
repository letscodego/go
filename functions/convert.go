package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	author := "nigel poulton"
	course := "getting started with kubernetes"

	fmt.Println(convert(author, course))
}

func convert(author, course string) (a, c string) {
	caser := cases.Title(language.English)
	author = strings.ToUpper(author)
	course = caser.String(course)

	return author, course
}
