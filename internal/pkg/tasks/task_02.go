package tasks

import (
	"fmt"
	"time"
)

func Run_task_02() {

	n := 6

	start := time.Now()

	fmt.Println(t02(n), time.Since(start))

}

func t02(n int) int {
	return fibRec(n)
}

func fibRec(n int) int {

	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	return fibRec(n-1) + fibRec(n-2)

}
