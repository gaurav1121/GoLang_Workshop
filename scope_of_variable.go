package main

import "fmt"

var word string = "Learning"

// https://play.golang.org/p/g610jbB1p-o
func main() {

	i := 12

	fmt.Println(word)
	fmt.Println(i)

	test()
}

func test() {

	// fmt.Println(i) //Error cause : i is not global defined
	fmt.Println(word) // word variable is global defined can be accessed anywhere
}
