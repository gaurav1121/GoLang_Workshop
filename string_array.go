package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	// Go compiler can catch indexing errors when constant is used
	// books[4] = "Neverland"

	// Go compiler cannot catch indexing errors when non-constant is used
	// i := 4
	// books[i] = "Neverland"

	// --------------------
	// PRINTING ARRAYS
	// --------------------

	// print the type
	fmt.Printf("books  : %T\n", books)

	// print the elements
	fmt.Println("books  :", books)

	// print the elements in quoted string
	fmt.Printf("books  : %q\n", books)

	// print the elements along with their types
	fmt.Printf("books  : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	// assign the first book as the wBook's first book
	wBooks[0] = books[0]

	// assign all the books starting from the 2nd book
	// to sBooks array

	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	// for i := 0; i < len(sBooks); i++ {
	// 	sBooks[i] = books[i+1]
	// }

	for i := range sBooks {
		sBooks[i] = books[i+1]
		// changes to sBooks[i] will not be visible here.
		// sBooks here is a copy of the original array.
	}
	// changes to sBooks are visible here

	// sBooks is a copy of the original sBooks array.
	//
	// v is also a copy of the original array element.
	//
	// if you want to update the original element, use it as in the loop above.
	//
	// for _, v := range sBooks {
	// 	v += "won't effect"
	// }

	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	// equal to: [4]bool
	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
