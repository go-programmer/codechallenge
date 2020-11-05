# Problem

Count the semiprime numbers in the given range [a..b] 

https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/

A prime is a positive integer X that has exactly two
distinct divisors: 1 and X. 
The first few prime integers are 2, 3, 5, 7, 11 and 13.

A semiprime is a natural number that is 
**the product of two (not necessarily distinct) prime numbers.**

The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

You are given two non-empty arrays P and Q, each consisting of M integers.
These arrays represent queries about the number of semiprimes within specified ranges.

Query K requires you to **find the number of semiprimes** 
within the range (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

For example, consider an integer N = 26 and arrays P, Q such that:
P[0] = 1    Q[0] = 26
P[1] = 4    Q[1] = 10
P[2] = 16   Q[2] = 20

The number of semiprimes within each of these ranges is as follows:
(1, 26) is 10,
(4, 10) is 4,
(16, 20) is 0.

Write a function: func Solution(N int, P []int, Q []int) []int

that, given an integer N and two non-empty arrays P and Q
consisting of M integers, returns an array consisting of 
M elements specifying the consecutive answers to all the 
queries.


For example, given an integer N = 26 and arrays P, Q such that:
    P[0] = 1    Q[0] = 26
    P[1] = 4    Q[1] = 10
    P[2] = 16   Q[2] = 20

the function should return the values [10, 4, 0], as explained above.


Write an efficient algorithm for the following assumptions:
N is an integer within the range [1..50,000];
M is an integer within the range [1..30,000];
each element of arrays P, Q is an integer within the range [1..N];
P[i] ≤ Q[i].


# Solution
### Analysis

**Semiprime: The product of two prime numbers.**

(1, 26) is 10 : 
1
2
3
4 2*2
5
6 2*3
7
8
9 3*3
10 2*5
11
12
13
14 2*7
15 3*5
16
17
18
19
20
21 3*7
22 2*11
23 
24 
25 5*5
26 2*13

(4, 10) is 4
4 2*2
5 
6 2*3  
7 
8 
9 3*3
10 2*5

(16, 20) is 0
16  2*2*2*2
17  1*17
18  2*3*3
19  1*19
20  2*2*5


From the lessons activities, primeFactors() returns
the number of smallets prime numbers, we can just 
count them. If 2 primes, increase semiprime.

### Pseudo count
Get the smallest prime numbers till N

Count the smallest prime numbers of all N

Count mark if the number is semiprime for all N

For each range, count the semi prime numbers in that range.
This logic can be improved counting the semi primes while counting
the semiprimes for all N but for now we do it the naive way.


The improved logic is using prefix sum.
But it gave wrong count for 4,10 = 3

(1, 26)  is 10
(4, 10)  is 4
(16, 20) is 0

(13, 20) is 2


(23, 29) is 2

When even include both ends, no substraction
When odd, substract max-min


### First solution
Scores:
Task Score 44%
Correctness 50%
Performance 40%

It failed small functional tests.

Failed small random:
WRONG ANSWER, got [17, 17, 16, 15, 10] expected [17, 17, 16, 15, 7] 


### Second solution

Checking solution based on the test cases.
> Min Prime Factor 60:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38  
0 0 0 0 2 0 2 0 2 3 2  0  2  0  2  3  2  0  2  0  2  3  2  0  2  5  2  3  2  0  2  0  2  3  2  5  2  0  2 

39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60
3  2  0  2  0  2  3  2  0  2  7  2  3  2  0  2  5  2  3  2  0  2


> Prefix Sum 40:
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38
0 0 0 1 1 2 2 2 3  4  4  4  4  5  6  6  6  6  6  6  7  8  8  8 9  10 10 10 10 10 10 10 11 12 13 13 13 14  

39 40
15 15


> Prefix sum 60:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38  
0 0 0 0 1 1 2 2 2 3  4  4  4  4  5  6  6  6  6  6  6  7  8  8  8  9 10 10 10 10 10 10 10 11 12 13 13 13 14 
                                       0  0  0  0  0  1  1  0  0  1 1 
                  1  1           1  1                 1  1        1 1                    1  1  1         1
39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60
15 15 15 15 15 15 15 16 16 16 17 17 18 18 18 18 19 19 20 21 21 21
1                    1        1  1  1           1     1  1  

(16, 27) 27-16 = 11(is Odd, max-min);  10 - 6 = 4; Verification: Correct

(10, 30) 20(even, include both ends, take max); PS[30] = 10; Verification: Wrong
PS[30]=10, PS[10]=4; Even if we do PS[30-10] = 6 
Prime Factors: 10, 13, 14, 15, 19, 20, 21 ; count = 7
Hence, the code logic had to be changed for even
`semiPrimes[i] = (prefixSumSemiPrime[max]` 
to 
`semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min]) + 1`

(8,31) 23 Odd, so max-min; 10-2 = 8; Verification: Correct
PFac: 9 10 14 15 21 22 25 26, count = 8

(15,18) 3 odd, max-min = 6-6 = 0; Verification: Wrong
PFac: 15, count = 1. 
In this case we might have to check if there was change from 14 to 15.
```
semiPrimes[i] = prefixSumSemiPrime[max] - prefixSumSemiPrime[min]
to
semiPrimes[i] = prefixSumSemiPrime[max] - prefixSumSemiPrime[min-1]
```

Gives same result 44.
The minor change flipped correctness of the previous result.
extreme_four fails, while small function pass, while 
small_functional still fails.

Performace tests gives same result

### Third submission
Even making the the minor changes for odd case (15,18) did not make
any difference to the score. It remaided as second.

### Fourth submission
Scores: Task Score 55% Correctness 75% Performance 40%

There was an issue when the smallest number was less than 4.
This caused an increase of 1 because the prime counts for 
numbers less than 4 are 0. Resolved the issue with a check
if the number was less than 4.
```go
else if (max-min)%2 == 0 {

    if prefixSumSemiPrime[max] == prefixSumSemiPrime[min] {
        semiPrimes[i] = 0

    } else {

        if min < 4 {
            semiPrimes[i] = prefixSumSemiPrime[max]

        } else {

            semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min]) + 1
        }
    }

} 
```

There are still failing cases:
WRONG ANSWER, 
got         [8, 13, __11__, 12, 9, 13, 8, 11,..
expected    [8, 13, __10__, 12, 9, 13, 8, 11,.. 

WRONG ANSWER,
got         [__49__, __51__, 46, 21, 42, 2.. 
expected    [__48__, __50__, 46, 21, 42, 2.. 

WRONG ANSWER, 
got         ..565, 6429, 6666, __1806__, 4139, __2416__, __6002__,..
expected    ..565, 6429, 6666, __1805__, 4139, __2415__, __6001__,.. 

WRONG ANSWER, 
got         ..955, 9085, 9244, `9878`, `10147`, 7624, 1062..
expected    ..955, 9085, 9244, `9877`, `10146`, 7624, 1062.. 

From the failing case analysis, there is just an increase of a number
in what is expected.

Checked it by adding similar condition for odd case
but it did not change the results.
```go
if min < 4 {
    semiPrimes[i] = prefixSumSemiPrime[max]

} else {
    semiPrimes[i] = prefixSumSemiPrime[max] - prefixSumSemiPrime[min-1]
}
```

To fix the issue, the problem was to find the 
test cases. As the lowest range was in 40, it
helped a lot to find that case.

(24, 30) previously was giving 3 semi prime count
which should have been 2. The issue was prefix sum
of semi prime included a value that change by previous
semiprime number.
Index                         : 22 23 24 25 26 27 28 28 30
SuffixSum of Semipeime numbers: 8  8  8  9  10 10 10 10 10

24 is not a semiprime number but in the suffix sum
is included which is incorrect. So the check was made
if there it really was a semiprime which is done by
comparing if previous number is lower than it. i.e,
the number to be semiprime, it has to be one more than
the previous. Here 23 also has 8, so 24 is not semiprime,
hence it should not be included in the total number of
semiprimes in that range.
```go
if prefixSumSemiPrime[min] > prefixSumSemiPrime[min-1] {
    semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min]) + 1

} else {
    semiPrimes[i] = (prefixSumSemiPrime[max] - prefixSumSemiPrime[min])
}
```

This solution scored 100.

The code coverage from the test is also 100 percent. 

There are alot of conditions which I believe can be improved.