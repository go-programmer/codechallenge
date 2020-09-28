
# Problem
An integer N is given, representing the area of some rectangle.

The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

The goal is to find the minimal perimeter of any rectangle whose area equals N. The sides of this rectangle should be only integers.

For example, given integer N = 30, rectangles of area 30 are:

        (1, 30), with a perimeter of 62,
        (2, 15), with a perimeter of 34,
        (3, 10), with a perimeter of 26,
        (5, 6), with a perimeter of 22.

Write a function:

    class Solution { public int solution(int N); }

that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

For example, given an integer N = 30, the function should return 22, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [1..1,000,000,000].


# Solution

```go
// Returns the number of divisors of n
// O(√n)
func solution(n int) int {
	i := 1
	result := 0

	for i*i < n {

		if n%i == 0 {
			result += 2
		}
		i++
	}

	if i*i == n {
		result++
	}

	return result
}
```
From the example above,
 (5, 6), with a perimeter of 22.

The perimeter is less when the two factorials
5,6, are relaively close to each other. As the 
max is 1Billion a probable number, we need to 
find a more efficient algorithm.

The straight forward naive solution would be 
calculate the factorial, calculate the perimeter.
The above method to find the factorial is O(N).

Thinking more about number, which can be odd or even,
composite or prime. 

Prime number, divisible by one and itself, so the 
perimeter will be 1 and itself. But we need to determine
whether that number is prime or not. 

All prime numbers are odd except 2. 
The number 2, which is the only even prime number.


After analyzing the divisor code further, it could
be understood that the larget divisor would be the 
last value of i, the difference of it's divident
would be the smallest value. 

Hence, the solution would be, 
- get the max divisor.
- get dividend
- get the perimeter.


The solution failed ✘: 
1 for runtime error

36 and 100,000,000 for wrong answer

36: got 26 expected 24
Max divisor is 4, which is a square number.
The dividend is 9, which also a square number.
Could it be that if divisor and divident are 
square numbers, we need to subatract 2?
Is this number special case?

9*2 + 4*2 = 18 + 8 = 26


100,000,000: got 41000 expected 40000 
This test was under the title square that
gave more hint on it issue.
It is a square number and the perimeter of
a square is 4*srqt(n).
sqrt(100,000,000) = 10000
4 * 10000 = 40000



Hence, the solution would be, 
- check if the number is 1

- check if the number is a perfect square.
  - return square root * 4

- get the max divisor.
- get dividend
- check if divisor and divident are both squares THEORY
  -  if both are square then reduce 1 from each
- get the perimeter.

The final solution scored 100 but the case of 36 
divisor and divident needed to be reduced to 1 is
not clear.