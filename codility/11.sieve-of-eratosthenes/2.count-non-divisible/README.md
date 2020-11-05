# Problem

Calculate the number of elements of an array that are not divisors of each element. 

https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible

You are given an array A consisting of N integers.

For each number A[i] such that 0 ≤ i < N, 
we want to **count the number of elements of the array that are not the divisors of A[i]**.
We say that these elements are non-divisors.

For example, consider integer N = 5 and array A such that:
A[0] = 3
A[1] = 1
A[2] = 2
A[3] = 3
A[4] = 6

For the following elements:
A[0] = 3, the non-divisors are: 2, 6,
A[1] = 1, the non-divisors are: 3, 2, 3, 6,
A[2] = 2, the non-divisors are: 3, 3, 6,
A[3] = 3, the non-divisors are: 2, 6,
A[4] = 6, there aren't any non-divisors.

Write a function: func Solution(A []int) []int

that, given an array A consisting of N integers, 
**returns a sequence of integers representing the amount of non-divisors.**

Result array should be returned as an array of integers.

For example, given:
    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6

the function should return [2, 4, 3, 2, 0], as explained above.

Write an efficient algorithm for the following assumptions:
N is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..2 * N].

# Solution

### Initial Solution
Sort. If sorted, we will not have to search top to bottom.
Just look for numbers smaller than it.

This solution became a bit hard to calculate 
total number of non divisors. And also duplicate
element was giving hard time to think clearly.

```golang
sort.Ints(A)
i := 1
for i < lenA {

    nonDivisorsCount := 0

    if i != lenA-1 {

        if copyA[i] != copyA[i+1] {

            //check if the number already exists in map
            if _, exists := mapA[copyA[i]]; !exists {

                mapA[copyA[i]] = 0

                // check previous numbers divide i
                for j := i - 1; j >= 0; j-- {

                    if copyA[i]%copyA[j] != 0 {
                        nonDivisorsCount++
                    }
                }

                mapA[copyA[i]] += lenA - (1 + i + nonDivisorsCount)
            }
        }
    }
    i++
}
```

### Second Solution 
Create a data struture to hold the elements 
and non divisors count.

```golang
type Element struct {
	index            []int
	count            int
	nonDivisorsCount int
}
```

Use map for the main calculation of non divisors.

It solved the problem and scored 66%, but scored 
100 on correctness. It failed large scale datasets
raising timeouts.

Time complexity is O(N^2). 

Created issue for optimization task at 
https://github.com/GoCheatsheet/codechallenge/issues/4


### Refactor/Optimization

The first solution would be more optimal 
but the problem was finding the divisors 
when the index grew.

We can store the divisors of the numbers
and for the next higher number can check
if the previous number divides it or not
and just add one to the divisors count 
and store the total number of divisors 
count. But the question is till which 
number do we trace back to find the divisors?

The concepts of Sieve needs to be looked 
at more to build a solution to find the 
divisors.


2020-11-5
Reviews the solution and found two solutions.
https://codesays.com/2014/solution-to-count-non-divisible-by-codility
https://www.martinkysel.com/codility-countnondivisible-solution


Martinkysel solution scored 100.
```python

Using the Sieve of Eratosthenes, 
you generate divisors for all input 
elements of A. If a given number x
is a divisor of element (x*N == element),
then N also is a divisor. 
(N = element/x). 

After all divisors are computed,
we simply subtract those 
(multiplied by their counts or 0)
from the total number of elements in A.

def solution(A):
  
    A_max = max(A)
  
    count = {}
    for element in A:
        if element not in count:
            count[element] = 1
        else:
            count[element] += 1
  
    divisors = {}
    for element in A:
        divisors[element] = set([1, element])
  
    # start the Sieve of Eratosthenes
    divisor = 2
    while divisor*divisor <= A_max:
        element_candidate = divisor
        while element_candidate  <= A_max:
            if element_candidate in divisors and not divisor in divisors[element_candidate]:
                divisors[element_candidate].add(divisor)
                divisors[element_candidate].add(element_candidate//divisor)
            element_candidate += divisor
        divisor += 1
  
    result = [0] * len(A)
    for idx, element in enumerate(A):
        result[idx] = (len(A)-sum([count.get(divisor,0) for divisor in divisors[element]]))
  
    return result
```
https://app.codility.com/demo/results/trainingW53CBG-MRS

