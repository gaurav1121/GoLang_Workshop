package main

import "fmt"

// https://play.golang.org/p/h_lp3NoDGj-
func main() {

	var s string
	s = "how are you?" //string
	fmt.Println(s)
	s = `how are you?` // string literal (raw string)
	fmt.Println(s)

	s = "<html> \n\t <body> \"hello\"</html>" //string
	fmt.Println(s)
	s = `
<html>
<body> "hello" </body>
</html> ` // raw string
	fmt.Println(s)
	fmt.Println("c:\\my\\dir\\file") //string
	fmt.Println(`c:\my\dir\file`)    //raw string
}
