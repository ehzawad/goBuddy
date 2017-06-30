package main

import (
	"fmt"
)

func main() {
	fmt.Printf("You win: ")

	doorNumber := 2

	switch doorNumber {
	case 1:
		fmt.Println("a new shit")

	case 2:
		fmt.Println("a new bullshit")

	default:
		fmt.Println("fuck")
	}

}
