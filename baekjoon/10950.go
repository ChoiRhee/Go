package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 1; i <= N; i++{
		var A, B int
		fmt.Scanln(&A, &B)
		fmt.Println(A + B)
	}
}