package main

import "fmt"

func main() {
	s := []int{11, 22, 33, 44}
	s = append(s, 5, 6)
	s = append(s, 10, 100)
	fmt.Println(s)
}
