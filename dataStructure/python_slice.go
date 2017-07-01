package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)
	firstSlice := a[0:2]
	secondSlice := a[1:4]
	fmt.Println(firstSlice, secondSlice)
	a[1] = 777
	fmt.Println(a, firstSlice, secondSlice)
	firstSlice = firstSlice[0:4]
	fmt.Println(a, firstSlice, secondSlice)

	fmt.Println("len(firstSlice):", len(firstSlice), "cap(firstSlice);", cap(firstSlice))
}
