package main

import (
	"fmt"
)

func main() {
	var a, b uint16 = 123, 456
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(^a)
}
