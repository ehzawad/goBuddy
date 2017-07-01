package main

import "fmt"

type Minutes int

func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}

func main() {
	minutes := Minutes(58)

	for i := 1; i <= 3; i++ {
		// (&minutes).Increment()
		// auto pointer type detection
		minutes.Increment()
		fmt.Println(minutes)
	}
}
