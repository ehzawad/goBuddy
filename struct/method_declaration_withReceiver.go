package main

import (
	"fmt"
	"strings"
)

type Title string

// recall uppercase method name so it will automatically exported
func (t Title) FixCase() string {
	return strings.Title(string(t))
}

func main() {
	name := Title("the matrix")
	fixed := name.FixCase()
	fmt.Println(fixed)
}

// A method is a function with a special receiver argument: https://golang.org/ref/spec#Method_declarations
// Allows you to define new behaviors for values of a type
// Method declaration looks just like an ordinary function declaration, except it has extra receiver parameter before the function name
// Why "receiver"? Think of calling a method on an object as sending it a message.
// No this or self, just an ordinary parameter name for the receiver

// The dot notation for calling methods reflects the way that the method is declared.
// Reciever occurs before method name, just like in the declaration

// Compiler looks at type of receiver, then calls method with that name defined on that type
// Don't confuse with package qualifier for exported function names; not the same thing
