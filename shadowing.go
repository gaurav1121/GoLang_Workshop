package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://play.golang.org/p/S5or0bckhoW
func main() {
	var n int

	if a := os.Args; len(a) != 2 {

		// only : a variable
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// only : a, n and err variable
		fmt.Printf("cannot convert %q. \n", a[1])

	} else {
		// all the variable in the if statements
		fmt.Printf("%s * 2 %d\n", a[1], n*2)
	}

	fmt.Printf("n is %d - You've been shadowed ;)\n", n)
}
