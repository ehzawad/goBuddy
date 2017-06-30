package main

import "fmt"

func main() {
	fmt.Println("You can't alter an existing pointer like you can in C or C++")

	var a int = 5
	var aPonter *int = &a

	fmt.Println("value: ", a)
	fmt.Println("address: ", aPonter)
	fmt.Println("addressed value: ", *aPonter)
}
