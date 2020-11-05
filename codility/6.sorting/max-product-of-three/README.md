# Problem

Maximize A[P] * A[Q] * A[R] for any triplet (P, Q, R).

A non-empty array A consisting of N integers is given. 
The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:
  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6

contains the following example triplets:

        (0, 1, 2), product is −3 * 1 * 2 = −6
        (1, 2, 4), product is 1 * 2 * 5 = 10
        (2, 4, 5), product is 2 * 5 * 6 = 60

Your goal is to find the maximal product of any triplet.

Write a function:

    func Solution(A []int) int

that, given a non-empty array A, returns the value of **the maximal product of any triplet**.

For example, given array A such that:
  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6

the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [3..100,000];
        each element of array A is an integer within the range [−1,000..1,000].


# Solution
Sort A
Get the product of of last 3.


The above solution works only for positive numbers.
It only achieved 44% as it failed for cases like:
-5, 5, -5, 4

After sort, -5 -5 4 5.

We need to check if the number first two numbers sign.
More simply multiply, and compare the first two and last
two values.

Check if first two product > second last * third last 
-5*-5 > -5 * 4, if true take first two else last 3 product.

Lets take another example:
-15 -5 -4 1 2 3
We can only take two smallest negative number
if their absolute values are higher, or simply their
multiplication as above.

Lets take another example:
-15 -5 -4 2 3


Sort A
Check if len is 3, return their product.
Compare first two and second and third last product.
Take the higher number and multiply with the last.


The second attempt failed on test case:
[-5, -6, -4, -7, -10], got -280 expected -120.

We need to check if all the answers are negative.

### Pseudocode:
If len is 3, return their product.

Sort A

If A[len-1] < 0, return the product of A[len-1],A[len-2],A[len-3]

Compare first two and second and third last product.
If A[0] * A[1] > A[len-2] * A[len-3]:
  return A[0] * A[1] * A[len-1]


Return the product of A[len-1],A[len-2],A[len-3]
