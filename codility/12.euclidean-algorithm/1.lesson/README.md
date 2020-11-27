# Theory Implementation

Euclidean Algorithm 
It solves the problem of computing the greatest common divisor (gcd) of two positive
integers.

Simple usecase: A 24-by-60 rectangle is covered with ten 12-by-12 square tiles, where 12 is the GCD of 24 and 60. More generally, an a-by-b rectangle can be covered with square tiles of side-length c only if c is a common divisor of a and b.

The least common multiple (lcm) of two integers
a and b is the smallest positive integer that
is divisible by both a and b.
lcm(a, b) = a·b / gcd(a,b)


### Found New Solution
Found this neat little solution while solving 
Google FooBar challenge.

https://github.com/oneshan/foobar/blob/master/bomb_baby/solution.py


```go
func gcdNew(x, y int) int {
	for y >= 1 {		
		fmt.Printf("%v : %v\n", x, y)

		if x < y {
			x, y = y, x
		}
		x, y = y, x % y
		
	}
	return x
}
```