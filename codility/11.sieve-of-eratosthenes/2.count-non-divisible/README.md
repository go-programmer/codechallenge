# Problem

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

Sort. If sorted, we will not have to search top to bottom.
Just look for numbers smaller than it.

Get factorial and store them in a map.
map[int][]int; key=number, factorial array.

A[0], map[1]{1}
A[1], map[2]{1,2}
A[2], map[3]{1,3}
A[3], map[6]{1,2,3,6}

Create a new ND array of length A. Fill 
ND index value with count from map.


Using the map solved the problem and scored 66%, 
100 on correctness but timed out on medium and 
large datasets.



Created issue for optimization task at
https://github.com/GoCheatsheet/codechallenge/issues/4




