package main

// :D :D :D :D
// https://gist.github.com/hasinhayder/0dd5c6ec13fbdd96183db36d79fefd78

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	letters := [][]string{{"e", "E"}, {"d"}, {" "}, {"mo", "mu"}, {"ba"}, {"ro", "ra", "roo"}, {"k"}}
	greetings := ""
	for _, chances := range letters {
		seeder := rand.NewSource(time.Now().UnixNano())
		randomizer := rand.New(seeder)
		greetings += chances[randomizer.Intn(len(chances))]
	}
	fmt.Println(greetings)
}
