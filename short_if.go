package main

import (
	"fmt"
	"strconv"
)

func main() {

	// https://play.golang.org/p/N9-J4xh9ed7
	// n, err := strconv.Atoi("42")

	// if err == nil {
	// 	fmt.Println("there was no error , n is ", n)
	// }

	// short if
	// https://play.golang.org/p/EdjNu5GYI38
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("there was no error , n is ", n)
	}
}
