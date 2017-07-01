package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan as return int
// channel chan int
func longTask(channel chan int) {
	fmt.Println("Starting long task")
	delay := rand.Intn(5) // delay is integer
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	channel <- delay
}

func main() {
	rand.Seed(time.Now().Unix())
	myChannel := make(chan int) // Add
	// go coroutine
	go longTask(myChannel)
	fmt.Println("Took", <-myChannel, "Second")
}
