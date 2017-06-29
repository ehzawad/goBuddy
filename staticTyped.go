package main

import (
	"fmt"
	"reflect"
)

// here type checked at compile time

func main() {
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1134134324322803))
	fmt.Println(reflect.TypeOf(1.1))
	fmt.Println(reflect.TypeOf("go"))
}
