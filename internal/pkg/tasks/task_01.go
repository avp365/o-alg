package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run_task_1() {
	a := 3.0
	n := 10.0

	start := time.Now()

	fmt.Println(T01_r(int(a), int(n)), time.Since(start))

	start = time.Now()

	fmt.Println(math.Pow(a, n), time.Since(start))
}

func T01(a, n int) int {

	answer := 1
	for i := 1; i <= n; i++ {
		answer = a * answer
	}

	return answer
}

func T01_r(a, n int) int {

	return powRec(a, n)
}

func powRec(a, n int) int {

	if n == 1 {
		return a
	}

	if n%2 == 0 {
		an := powRec(a, n/2)
		return an * an
	} else {
		an := powRec(a, n-1)
		return a * an
	}

}
