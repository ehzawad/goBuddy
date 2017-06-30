package main

import "fmt"

func main() {
	var a int
	a = 1
	var b, c int
	b = 2
	c = 3
	// this will automatically detect the type
	var d = 5

	// this is the most unique short variable declaration
	e := 6

	f, g := 7, 8

	fmt.Println(a, b, c, d, e, f, g)
}
