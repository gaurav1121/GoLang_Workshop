package main

import (
	"fmt"
	"path"
)

// https://play.golang.org/p/Ga4wg9jnhGd
func main() {
	var file string
	_, file = path.Split("css/main.css") // _ is a blank identifer (unused variable not allowed in go)
	fmt.Println("file", file)
}
