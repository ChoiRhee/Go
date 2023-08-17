package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var tmp float32
	var avgarr []int

	/// n개의 입력 값을 1000 곱해서 int형으로 arr에 추가
	for i := 0; i < n; i++ {
		fmt.Scanln(&tmp)
		avgarr = append(avgarr, int(tmp * 1000))
	}
	fmt.Println(avgarr)

	// 소수점 세자리까지니까 최대 사람 수는 1000명
	for peoplecnt := 1; peoplecnt <= 1000; peoplecnt++ {
		// 1부터 1000까지 인원수를 대입해서 맞는지 계산
		if (possiblecnt(peoplecnt, avgarr)) {
			fmt.Println(peoplecnt)
			break
		}
	}
}

func possiblecnt(peoplecnt int, avgarr []int) bool {
	for _, avg := range avgarr {
		var left int = 0
		var right int = 10 * peoplecnt
		var ispossible bool = false

		for left <= right {
			scoresum := (left + right) / 2
			currentavg := (scoresum * 1000) / peoplecnt
			if currentavg == avg {
				if currentavg > 10 * 1000 {
					continue
				}
				ispossible = true
				break
			} else if currentavg > avg {
				right = scoresum - 1
			} else {
				left = scoresum + 1
			}
		}
		if (!ispossible) {
			return false
		}
	}
	return true
}