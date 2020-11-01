
# Problem
Count the number of distinct slices (containing only unique numbers). 

An integer M and a non-empty array A consisting of N non-negative integers are given. 
All integers in array A are less than or equal to M.

A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. 
The slice consists of the elements A[P], A[P + 1], ..., A[Q]. 

A distinct slice is a slice consisting of only unique numbers. 
That is, no individual number occurs more than once in the slice.

For example, consider integer M = 6 and array A such that:
    A[0] = 3
    A[1] = 4
    A[2] = 5
    A[3] = 5
    A[4] = 2

There are exactly nine distinct slices: (0, 0), (0, 1), (0, 2), (1, 1), (1, 2), (2, 2), (3, 3), (3, 4) and (4, 4).

The goal is to calculate the number of distinct slices.

Write a function:

    func Solution(M int, A []int) int

that, given an integer M and a non-empty array A consisting of N integers, returns the number of distinct slices.

If the number of distinct slices is greater than 1,000,000,000, the function should return 1,000,000,000.

For example, given integer M = 6 and array A such that:
    A[0] = 3
    A[1] = 4
    A[2] = 5
    A[3] = 5
    A[4] = 2

the function should return 9, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [1..100,000];
        M is an integer within the range [0..100,000];
        each element of array A is an integer within the range [0..M].

# Solution
Although the problem mentions - "Count the number of distinct slices (containing only unique numbers)",
But the provided example solution includes a pair of 5,5 at indexes (2, 2), (3, 3).
So there is a duplicate.

The pair (2,3) which is (5,5) is excluded.


### Algorithm
Use map[int]int to store unique pairs.

The problem is there can be duplicates while 
making a pair of same index, (2, 2), (3, 3).
We have to check for that.


for i=0; i<lenA; i++
  keyNum=A[i]
  Store map[keyNum]=keyNum

  for j=i+1; j<lenA; j++
    Check if [keyNum]=keyNum exists in store
      break

    store [keyNum]= A[j]

The logic using map fails as the value of a
key is replaced. 

https://app.codility.com/demo/results/trainingPS2B5F-7N5

The logic was slightly modified and set was used.
That solution scored 100% in correctness and scored
40 in preformance.


Minor changes does not change the performace result.
```go
if lenA >= 1000000000 {
		return 1000000000
}
```

https://app.codility.com/demo/results/trainingTPU8FK-VAN

Found two solutions that gives 100% with O(N) time complexity:
https://www.martinkysel.com/codility-countdistinctslices-solution
https://codesays.com/2014/solution-to-count-distinct-slices-by-codility


```python
# https://app.codility.com/demo/results/trainingNT8D47-RNZ
def solution(M, A):
    the_sum = 0
    front = back = 0
    seen = [False] * (M+1)
    while (front < len(A) and back < len(A)):
        while (front < len(A) and seen[A[front]] != True):
            the_sum += (front-back+1)
            seen[A[front]] = True
            front += 1
        else:
            while front < len(A) and back < len(A) and A[back] != A[front]:
                seen[A[back]] = False
                back += 1
                 
            seen[A[back]] = False
            back += 1
    return min(the_sum, 1000000000)    
```

```python
def solution(M, A):
    accessed = [-1] * (M + 1)   # -1: not accessed before
                                # Non-negative: the previous occurrence position
    front, back = 0, 0
    result = 0
    for front in range(len(A)):
        if accessed[A[front]] == -1:
            # Met with a new unique item
            accessed[A[front]] = front
        else:
            # Met with a duplicate item.
            # Compute the number of distinct slices 
            # between newBack-1 and back position.
            newBack = accessed[A[front]] + 1
            result += (newBack - back) * (front - back + front - newBack + 1) / 2

            if result >= 1000000000:   return 1000000000

            # Restore and set the accessed array
            for index in range(back, newBack):
                accessed[A[index]] = -1
                
            accessed[A[front]] = front
            back = newBack

    # Process the last slices
    result += (front - back + 1) * (front - back + 2) / 2
    return min(int(result), 1000000000)
```