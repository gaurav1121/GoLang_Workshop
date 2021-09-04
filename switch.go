package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {

	case "Chandigarh":
		fmt.Println("India")
	case "tokyo":
		fmt.Println("japan")
	default:
		fmt.Println("where?")
	}

}
