package main

import "fmt"

func main() {

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)

	fmt.Println()
}

// https://play.golang.org/p/gBBRQmbu3F3
