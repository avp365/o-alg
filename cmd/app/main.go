package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	n := 4
	summ := 0

	data := calc(n, n)

	for _, valueInt := range data {
		summ += valueInt * valueInt
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println(summ, elapsed)

}

func calc(n int, ni int) []int {

	data := make([][]int, 9+1)
	res := make([]int, ni*9+1)
	resPrev := make([]int, ni*9+1)

	for i := range data {
		data[i] = make([]int, ni*9+1)
	}

	if ni > 1 {
		resPrev = calc(n, ni-1)

	}

	if ni == 1 {
		for s := 0; s <= ni*9; s++ {

			for i := 0; i <= 9; i++ {

				if s == i {
					data[s][i] = 1
				}
				continue
			}

		}
	}

	if ni > 1 {

		for i := 0; i <= 9; i++ {
			for s, v := range resPrev {
				data[i][i+s] = v
			}
		}

	}

	for s := 0; s <= ni*9; s++ {
		for i := 0; i <= 9; i++ {

			res[s] = res[s] + data[i][s]

		}

	}
	fmt.Println(res)
	return res

}
