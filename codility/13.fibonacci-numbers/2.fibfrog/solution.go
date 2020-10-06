package solution

func fibonacciDynamic(n int) int {
	fib := make([]int, n+2)
	fib[1] = 1

	i := 2
	for i < n+1 {
		fib[i] = fib[i-1] + fib[i-2]
		i++
	}

	return fib[n]
}
