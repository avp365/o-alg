package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run_task_11() {
	a := 3.0
	n := 10.0

	start := time.Now()

	fmt.Println(t11(int(a), int(n)), time.Since(start))

	start = time.Now()

	fmt.Println(math.Pow(a, n), time.Since(start))
}

func T11(a, n int) int {

	answer := 1

	if n%2 == 0 {
		for i := 1; i <= n/2; i++ {
			answer = a * answer
		}
		answer = answer * answer
	} else {
		for i := 1; i <= (n-1)/2; i++ {
			answer = a * answer
		}
		answer = a * answer * answer
	}

	return answer
}
