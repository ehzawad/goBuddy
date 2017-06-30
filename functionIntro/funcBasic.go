package main

import "fmt"

func main() {
	myFunc()
	fmt.Println(sum(3, 5))
	fmt.Println(diff(6, 2))
}

func myFunc() {
	fmt.Println("myFunc")
}

func sum(a int, b int) int {
	return a + b
}

func diff(big int, small int) (result int) {
	result = big - small
	return
}
