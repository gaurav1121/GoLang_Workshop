package main

import (
	"fmt"
	"os"
)

func main() {
	// https://play.golang.org/p/OfsK2gDYOhF
	// os.Args is a string Array ([]strings) where 0th index is file path
	// fmt.Printf("%#v\n", os.Args)

	// https://play.golang.org/p/uZNvA67svHV
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st Argument: ", os.Args[1])
	fmt.Println("2nd Argument: ", os.Args[2])
	fmt.Println("3rd Argument: ", os.Args[3])
	fmt.Println("No. of Items Inside os.Args: ", len(os.Args))
}
