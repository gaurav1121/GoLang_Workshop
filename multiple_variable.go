package main

import (
	"fmt"
)

// https://play.golang.org/p/ggwR6KkorBp
func main() {

	var i, j, k = 1, 2, 3
	var (
		name       = "Sangam Biradar"
		occupation = "Developer Advocate"
	)

	fmt.Println(i, j, k)
	fmt.Printf("%s is a %s\n", name, occupation)
}
