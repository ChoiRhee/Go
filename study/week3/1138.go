package main

import "fmt"

func main() {
	var n, tmp int
	fmt.Scanln(&n)
	
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, 0)
	}

	for i := 1; i < n + 1; i++ {
		fmt.Scanf("%d", &tmp)
		cnt := 0

		for j := 0; j <= n; j++ {
			if tmp == cnt && arr[j] == 0 { 
				arr[j] = i
				break
			} else if arr[j] == 0 {
				cnt += 1
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}