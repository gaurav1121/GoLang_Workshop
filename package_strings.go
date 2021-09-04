package main

import (
	"fmt"
	"os"
	"strings"
)

// https://play.golang.org/p/R4dJ8eWGr2Z
// https://play.golang.org/p/XNstN9cpu7Y
func main() {

	msg := os.Args[1]
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	fmt.Println(s)
	s = strings.ToUpper(s)
	fmt.Println(s)

}
