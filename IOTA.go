package main

import "fmt"

// Built-in constant number generator
func main() {
	// https://play.golang.org/p/2pqUCVqQc-3
	// const (
	// 	monday    = 0
	// 	tuesday   = 1
	// 	wednesday = 2
	// 	thursday  = 3
	// 	friday    = 4
	// 	saturday  = 5
	// 	sunday    = 6
	// )

	// https://play.golang.org/p/p6KTe_eg7U1
	const (
		// monday = iota
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, saturday, sunday)
}
