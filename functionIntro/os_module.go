package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, error := os.Stat("os_module.go")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(fileInfo.Size())
	}

	fileInfo2, error2 := os.Stat("hi.php")
	if error2 != nil {
		fmt.Println(error2)
	} else {
		fmt.Println(fileInfo2.Size())
	}
}
