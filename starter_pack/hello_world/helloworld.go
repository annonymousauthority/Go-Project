package main

import (
	"fmt"
	"strings"

	"rsc.io/quote"
)

/*
As a ritual nothing more.
*/

func main() {
	stringusage()
	arrays()
	quoteUsage()
}

func quoteUsage() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}

func arrays() {
	welcomeArray := [3]string{"Hello", "Sam", "Welcome"}
	fmt.Println(welcomeArray[0], welcomeArray[1], welcomeArray[2])
}

func stringusage() {
	friend := "Sam"
	fmt.Println("I was just calling", strings.ToUpper(friend))
	fmt.Println(strings.Count("free", "e"))
	fmt.Println(len(friend))
}
