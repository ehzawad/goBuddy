package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	square, err := squareRoot(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(square)

	q := divisionTwist(15, 0)

	fmt.Println(q)

}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("no square root of negative number")
	}
	return math.Sqrt(x), nil
}

func divisionTwist(divident float64, divisor float64) float64 {
	return divident / divisor
}
