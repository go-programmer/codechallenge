# Probolem
Check whether two numbers have the same prime divisors. 

A prime is a positive integer X that has exactly
two distinct divisors: 1 and X. The first few prime
integers are 2, 3, 5, 7, 11 and 13.

A prime D is called a prime divisor of a positive
integer P if there exists a positive integer K such
that D * K = P. For example, 2 and 5 are prime divisors of 20.

You are given two positive integers N and M. 
The goal is to check whether the sets of prime 
divisors of integers N and M are exactly the same.


For example, given:
N = 15 and M = 75, the prime divisors are the same: {3, 5};
N = 10 and M = 30, the prime divisors aren't the same: {2, 5} is not equal to {2, 3, 5};
N = 9 and M = 5, the prime divisors aren't the same: {3} is not equal to {5}.

Write a function: func Solution(A []int, B []int) int

that, given two non-empty arrays A and B of Z integers,
returns the number of positions K for which the prime
divisors of A[K] and B[K] are exactly the same.

For example, given:
    A[0] = 15   B[0] = 75
    A[1] = 10   B[1] = 30
    A[2] = 3    B[2] = 5

the function should return 1, 
because only one pair (15, 75) 
has the same set of prime divisors.

Write an efficient algorithm for the following assumptions:
Z is an integer within the range [1..6,000];
each element of arrays A, B is an integer within the range [1..2,147,483,647].


# Solution

A direct or naive way to solve would be to take
the prime factorials of the number and compare
but this would not use the lessons concepts, GCD, LCM.


A[0] = 15   B[0] = 75 
75 mod 15 = 0
GCD=15 
LCM=15*75 / 15 = 75


A[1] = 10   B[1] = 30 
30 mod 10 = 0
GCD = 10
LCM = 10*30 / 10 = 30


A[2] = 3    B[2] = 5
5%3 != 0
Both are prime numbers, which will not have any common divisors.

N = 15 and M = 75, the prime divisors are the same: {3, 5};
N = 10 and M = 30, the prime divisors aren't the same: {2, 5} is not equal to {2, 3, 5};
N = 9 and M = 5, the prime divisors aren't the same: {3} is not equal to {5}.   


### Second Thought 
Cheking via comparing prime factors would be quadratic in nature,
I was looking for something that uses GCD implementation and 
found a thirdparty solution. 

The basic idea was that for the number to have equal number of
prime factors means it should be divisible equally by its GCD.
Or the number has a remainder 1 or not.


