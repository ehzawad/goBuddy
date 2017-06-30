package main

import "fmt"

func main() {
	CountByThrees(3, 9)
}

func CountByThrees(start int, end int) {
	// YOUR CODE HERE
	for i := start; i <= end; i += 3 {
		fmt.Println(i)
	}

	if false || true {
		fmt.Println("true") // printed
	}
}
