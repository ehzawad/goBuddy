package main

import "fmt"

func main() {
	var months [3]string
	months[0] = "jan"
	months[1] = "feb"
	months[2] = "mar"

	fmt.Println(months[0])

	for index, month := range months {
		fmt.Println(index, "-->", month)
	}

	// when you don't need indexing
	for _, month := range months {
		fmt.Println(month)
	}
}
