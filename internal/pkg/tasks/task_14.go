package tasks

import (
	"fmt"
	"math"
	"time"
)

func Run_task_14() {

	n := 8

	start := time.Now()

	fmt.Println(t14(n), time.Since(start))

}

func t14(n int) int {
	return goldenRatio(n)
}

func goldenRatio(n int) int {

	φ := (1 + math.Sqrt(5)) / 2
	F := math.Pow(φ, float64(n)) / math.Sqrt(5)

	return int(F)
}
