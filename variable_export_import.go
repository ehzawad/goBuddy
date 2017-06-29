package main

import "fmt"

// package-level and can be exported
var Hell = "Lucifer"
var Number = 99

func main() {

	// if the first letter is lower-case, the variable is unexported and can only be used within the current package(THE SAME IS TRUE FOR CONSTANTS, FUNCTIONS etc)

	myHell := "THis is a hell"
	myNum := 5

	fmt.Println(myHell, myNum)

	// if the first letter is upper-case, the varibale is exported and can also be used OUTSIDE the current package
}
