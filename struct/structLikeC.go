package main

import "fmt"

type Person struct {
	Name    string
	Versity string
	age     float64
}

func main() {
	guy := Person{}
	guy.Name = "Abrar Ghalib"
	guy.Versity = "AIUB"
	guy.age = 22

	fmt.Println(guy)
}
