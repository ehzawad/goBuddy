package main

import "fmt"

func main() {
	ages := map[string]float64{}

	ages["Alice"] = 12
	ages["Bob"] = 9

	fmt.Println(ages)
	fmt.Println(ages["Alice"])
	delete(ages, "Bob")
	fmt.Println(ages)

	friends := map[string]string{"Habib": "ProblemSolver", "Abrar": "VRdeveloper"}
	fmt.Println(friends)
	fmt.Println(friends["Abrar"])

	// key value pairs
	for name, hobby := range friends {
		fmt.Println(name, hobby)
	}

	// only key
	for name := range friends {
		fmt.Println(name)
	}

	// only value, _ is used for omitting key
	for _, hobby := range friends {
		fmt.Println(hobby)
	}
}
