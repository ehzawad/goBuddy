package main

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

func main() {
	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(reflect.TypeOf(time.Now()))
}
