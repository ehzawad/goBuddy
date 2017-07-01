package main

import "fmt"

type Minutes int
type Hours int
type Weight float64
type Title string
type Answer bool

func main() {
	minutes := Minutes(37)
	minutes += 3
	hours := Hours(3)
	weight := Weight(22.2)
	name := Title("Zawad")
	answer := Answer(true)
	fmt.Println(minutes, hours, weight, name, answer)

	if weight > Weight(10) {
		fmt.Println("Adult")
	}

	// but this will not work
	// though int
	// their custom type is mismatched
	// if Minutes(10) > Hours(1) {
	// 	Println("Shit")
	// }
}
