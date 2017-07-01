package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 100}
	for _, v := range s {
		fmt.Println(v)
	}
}
