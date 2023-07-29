package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	if a == b && b == c {
		fmt.Println(10000 + a * 1000)
	} else if a == b || a == c {
		fmt.Println(1000 + a * 100)
	} else if b == c {
		fmt.Println(1000 + b * 100)
	} else {
		fmt.Println(math.Max(math.Max(a, b), c) * 100)
	}
}