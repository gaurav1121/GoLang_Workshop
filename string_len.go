package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// https://play.golang.org/p/vNU-AYDSEm7
	name := "Gaurav"
	fmt.Println(len(name))

	nname := "गौरव"
	fmt.Println(
		utf8.RuneCountInString(nname))
}
