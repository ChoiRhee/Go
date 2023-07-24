package main

import "fmt"

func main() {
	var s, sum, count int
	fmt.Scanln(&s)
	
	for i := 1; i <= s; i++ {
		sum += i
		count += 1
		if sum > s {
			fmt.Println(count-1)
			break
		} else if sum == s {
			fmt.Println(count)
			break
		}
		
	}
}