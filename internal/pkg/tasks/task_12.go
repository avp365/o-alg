package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run_task_12() {
	a := 3.0
	n := 10.0

	start := time.Now()

	fmt.Println(t11(int(a), int(n)), time.Since(start))

	start = time.Now()

	fmt.Println(math.Pow(a, n), time.Since(start))
}

func t12(a, n int) int {

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
