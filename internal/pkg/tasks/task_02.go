package tasks

import (
	"fmt"
	"time"
)

func Run_task_02() {

	n := 8

	start := time.Now()

	fmt.Println(T02(n), time.Since(start))

}

func T02(n int) int {
	return fibIterN(n)
}

func fibRecN2(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibRecN2(n-1) + fibRecN2(n-2)

}

func fibIterN(n int) int {
	in_2 := 0
	in_1 := 1
	in := 0

	for ni := 2; ni <= n; ni++ {

		in = in_1 + in_2
		in_2 = in_1
		in_1 = in
	}

	return in
}
