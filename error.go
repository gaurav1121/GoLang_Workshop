package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// https://play.golang.org/p/YCqwRIcQGuE
	// n, error := strconv.Itoa(os.Args[1])
	// fmt.Println("Converted Number :", n)
	// fmt.Println("Returned error value :", error)

	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Sucess: Converted %q to %d \n", age, n)

}
