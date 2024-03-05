package task

func T1(a, n int) int {

	answer := 1
	for i := 1; i <= n; i++ {
		answer = a * answer
	}

	return answer
}
