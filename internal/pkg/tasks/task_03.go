package tasks

import (
	"fmt"
	"time"
)

func Run_task_03() {

	n := 115

	start := time.Now()

	fmt.Println(t03(n), time.Since(start))

}

func t03(n int) int {
	return f01(n)

}

func f01(n int) int {

	count := 0

	for i := 2; i <= n; i++ {

		if isPrime(i) {
			count++
		}

	}

	return count
}

func isPrime(n int) bool {
	count := 0

	for i := 1; i <= n; i++ {

		if n%i == 0 {
			count++
		}
	}
	return count == 2
}
