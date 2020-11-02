# Problem
Find the minimal absolute value of a sum of two elements.



Let A be a non-empty array consisting of N integers.

The abs sum of two for a pair of indices (P, Q) is the absolute value |A[P] + A[Q]|, for 0 ≤ P ≤ Q < N.

For example, the following array A:
  A[0] =  1
  A[1] =  4
  A[2] = -3

has pairs of indices (0, 0), (0, 1), (0, 2), (1, 1), (1, 2), (2, 2).
The abs sum of two for the pair (0, 0) is A[0] + A[0] = |1 + 1| = 2.
The abs sum of two for the pair (0, 1) is A[0] + A[1] = |1 + 4| = 5.
The abs sum of two for the pair (0, 2) is A[0] + A[2] = |1 + (−3)| = 2.
The abs sum of two for the pair (1, 1) is A[1] + A[1] = |4 + 4| = 8.
The abs sum of two for the pair (1, 2) is A[1] + A[2] = |4 + (−3)| = 1.
The abs sum of two for the pair (2, 2) is A[2] + A[2] = |(−3) + (−3)| = 6.

Write a function:

    func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the minimal abs sum of two for any pair of indices in this array.

For example, given the following array A:
  A[0] =  1
  A[1] =  4
  A[2] = -3

the function should return 1, as explained above.

Given array A:
  A[0] = -8
  A[1] =  4
  A[2] =  5
  A[3] =-10
  A[4] =  3

the function should return |(−8) + 5| = 3.

Write an efficient algorithm for the following assumptions:
N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].


# Solution

The general rough idea would to iterate with two loops,
one outer to traverse all and second to get the sums.
So this solution time constraint would be O(N^2).

The general idea of caterpillar method seems to be to
use two pointers, front and back, and I searched how
others solved the problem.

Found a solution by codesays.com
https://codesays.com/2014/solution-to-min-abs-sum-of-two-by-codility

```python
def solution(A):
    A.sort()        # Sort A in non-decreasing order
    if A[0] >= 0:   return A[0] + A[0]      # All non-negative
    if A[-1] <= 0:  return -A[-1] - A[-1]   # All non-positive
    front, back = len(A)-1, 0
    minAbs = A[-1] + A[-1]                  # Final result
    # Travel the array from both ends to some center point.
    # See following post for the proof of this method.
    # https://codesays.com/2014/solution-to-min-abs-sum-of-two-by-codility
    while back <= front:
        temp = abs(A[back] + A[front])
        # Update the result if needed
        if temp < minAbs:  minAbs = temp
        
        # Adjust the pointer for next trying
        if back == front: 
            break
        elif abs(A[back+1] + A[front]) <= temp:     
            back += 1
        elif abs(A[back] + A[front-1]) <= temp:     
            front -= 1
        else:                                       
            back += 1
            front -= 1
    return minAbs
```

This solution scored 100 and with a time complexity of O(N * log(N)) 

https://app.codility.com/demo/results/trainingQ7EJQB-7D7


