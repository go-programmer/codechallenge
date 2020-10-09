
# Problem

You are given integers K, M and a non-empty array A consisting of N integers.
Every element of the array is not greater than M.

You should **divide this array into K blocks of consecutive elements**.
The size of the block is any integer between 0 and N.
Every element of the array should belong to some block.

The sum of the block from X to Y equals A[X] + A[X + 1] + ... + A[Y].
The sum of empty block equals 0.

The large sum is the maximal sum of any block.

For example, you are given integers K = 3, M = 5 and array A such that:
  A[0] = 2
  A[1] = 1
  A[2] = 5
  A[3] = 1
  A[4] = 2
  A[5] = 2
  A[6] = 2

The array can be divided, for example, into the following blocks:
[2, 1, 5, 1, 2, 2, 2], [], [] with a large sum of 15;
[2], [1, 5, 1, 2], [2, 2] with a large sum of 9;
[2, 1, 5], [], [1, 2, 2, 2] with a large sum of 8;
[2, 1], [5, 1], [2, 2, 2] with a large sum of 6.

The goal is to **minimize the large sum**. 
In the above example, 6 is the minimal large sum.

Write a function:func Solution(K int, M int, A []int) int

that, given integers K, M and a non-empty array A consisting of N integers, returns the minimal large sum.

For example, given K = 3, M = 5 and array A such that:
  A[0] = 2
  A[1] = 1
  A[2] = 5
  A[3] = 1
  A[4] = 2
  A[5] = 2
  A[6] = 2

the function should return 6, as explained above.

Write an efficient algorithm for the following assumptions:
N and K are integers within the range [1..100,000];
M is an integer within the range [0..10,000];
each element of array A is an integer within the range [0..M].

# Solution
### Analysis
[2, 1, 5, 1, 2, 2, 2], [], [] with a large sum of 15;
[2], [1, 5, 1, 2], [2, 2] with a large sum of 9;
[2, 1, 5], [], [1, 2, 2, 2] with a large sum of 8;
[2, 1], [5, 1], [2, 2, 2] with a large sum of 6.

Create k blocks of array of any sizes.
The largest sum should be the smallest compared to
the other boxes created from the combination of 
the elements of the array.

0 1 2 3 4 5 6
2 1 5 1 2 2 2

The minimum largest sum value that the array can 
have is 5. So probably it will be a good idea to
make it as a minimum bench mark.

[2 1] [5] [1 2 2 2]
When a sinlge element array with value 5 is taken,
the maximum sum is 1+2+2+2 = 7.

Or, When we take a single integer element with value 1,
[2] [1] [5 1 2 2 2], the max sum is 5+1+2+2+2=12

Lets suppose K = 4,
[2 1] [5] [1 2] [2 2]
Then the minimum block of sum 5 would be possible.

What happens when we bring two max numbers, two 5?
[2 1] [5 1] [2 2 2] [5] 
Then the min sum would be 6.

Supppose second 5 is not at the end.
[2] [1 5 1] [2 2 2] [5 3] the max will be 8


This seems like a Dynamic Technique related problem.


With some google, found two resources
https://codesays.com/2014/solution-to-min-max-division-by-codility
https://www.martinkysel.com/codility-minmaxdivision-solution

Both of them solved it in same manner.
```python


def blockSizeIsValid(A, max_block_cnt, max_block_size):
    block_sum = 0
    block_cnt = 0
 
    for element in A:
        if block_sum + element > max_block_size:
            block_sum = element
            block_cnt += 1
        else:
            block_sum += element
        if block_cnt >= max_block_cnt:
            return False
 
    return True
 
def binarySearch(A, max_block_cnt, using_M_will_give_you_wrong_results):
    lower_bound = max(A)
    upper_bound = sum(A)
 
    if max_block_cnt == 1:      return upper_bound
    if max_block_cnt >= len(A): return lower_bound
 
    while lower_bound <= upper_bound:
        candidate_mid = (lower_bound + upper_bound) // 2
        
        if blockSizeIsValid(A, max_block_cnt, candidate_mid):
            upper_bound = candidate_mid - 1
        else:
            lower_bound = candidate_mid + 1
 
    return lower_bound
 
def solution(K, M, A):
    return binarySearch(A,K,M)
```

Converted the solution in Go but the Go solution gives
wrong answer. 

```python
candidate_mid = (lower_bound + upper_bound) // 2
```
The use of `//` might be giving a different results.
`//` is the floor division operator. It produces 
the floor of the quotient of its operands, without
floating-point rounding for integer operands. 

https://stackoverflow.com/questions/14444520/two-forward-slashes-in-python

And exact representation in Go was not available.
So I tried to convert into float and take the floor
value. But it did now work. Checking furhter.

What Every Computer Scientist Should Know About Floating-Point Arithmetic
https://docs.oracle.com/cd/E19957-01/806-3568/ncg_goldberg.html

### Failure Analysis
Review mid  and return value
mid = 10 | 10
ret = true | false

mid = 7
ret=true

mid=5 
ret=false

mid=6
ret=true


After the analysis, I have found the mistake to be silly
```go
for element := range A  // Bug
for _, element := range A // Resolved
```

When the Python code was reengineered to Go,
`range` implementation got wrong. range in 
Go return index, where as value is returned
in Python.

The solution works fine.
