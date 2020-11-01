
# Problem
Count the number of triangles that can be built from a given set of edges. 


An array A consisting of N integers is given. 
A triplet (P, Q, R) is triangular if it is 
possible to build a triangle with sides of
lengths A[P], A[Q] and A[R]. 

In other words, triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:
A[P] + A[Q] > A[R],
A[Q] + A[R] > A[P],
A[R] + A[P] > A[Q].

For example, consider array A such that:
  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 12

There are four triangular triplets that
can be constructed from elements of this
array, namely (0, 2, 4), (0, 2, 5), (0, 4, 5), and (2, 4, 5).

Write a function:
    func Solution(A []int) int

that, given an array A consisting of N integers,
returns the number of triangular triplets in this array.

For example, given array A such that:
  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 12

the function should return 4, as explained above.

Write an efficient algorithm for the following assumptions:
N is an integer within the range [0..1,000];
each element of array A is an integer within the range [1..1,000,000,000].


# Solution 
Found two solutions that scored 100

https://www.martinkysel.com/codility-counttriangles-solution
```python
# Time complexity O(N^2)
def solution(A):
    A.sort()
    # print A
    triangle_cnt = 0
     
    for P_idx in range(0, len(A)-2):
        Q_idx = P_idx + 1
        R_idx = P_idx + 2
        
        while (R_idx < len(A)):
            if A[P_idx] + A[Q_idx] > A[R_idx]:
                triangle_cnt += R_idx - Q_idx
                R_idx += 1
            elif Q_idx < R_idx -1:
                    Q_idx += 1
            else:
                R_idx += 1
                Q_idx += 1
                 
    return triangle_cnt

# https://app.codility.com/demo/results/trainingWVQPKC-R4A
```

https://codesays.com/2014/solution-to-count-triangles-by-codility

```python
# Time complexity O(N^2) 
def solution(A):
    n = len(A)
    result = 0
    A.sort()
    for first in range(n-2):
        third = first + 2
        for second in range(first+1, n-1):
            while third < n and A[first] + A[second] > A[third]:
                third += 1
            result += third - second - 1
    return result

# https://app.codility.com/demo/results/trainingAT4JBT-4TE/

```

Codesays solution seems to be more elegant.

