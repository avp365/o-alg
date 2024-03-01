package main

import (
	"fmt"
	"math"
	"time"

	"github.com/avp365/arch-pat/internal/pkg/task"
)

func main() {
	a := 3.0
	n := 5.0
	task.T1(int(a), int(n))
	start := time.Now()

	answer := math.Pow(a, n)

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(answer, elapsed)
}
