# Problem
Rotate an array to the right by a given number of steps.


# Solution
The solution to the problem was found and resoloved.

https://app.codility.com/demo/results/training9HH2PD-PDF

```go
func solution(A []int, K int) []int {
	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A
	}

	if len(A) < K {
		K = K % len(A)
	}

	lhs := A[len(A)-K:]

	return append(lhs, A[:len(A)-K]...)
}
```

Trying out left rotation to understand the implementation better.

