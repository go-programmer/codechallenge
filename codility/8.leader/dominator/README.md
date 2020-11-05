# Problem:
Find an index of an array such that its value occurs at more than half of indices in the array. 

An array A consisting of N integers is given.

The dominator of array A is the value **that occurs in more than half** of the elements of A.

For example, consider array A such that
 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3

The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

Write a function

    func Solution(A []int) int

that, given an array A consisting of N integers,
**returns index of any element of array A** in which the dominator of A occurs.

The function should return −1 if array A does not have a dominator.

For example, given array A such that
 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3

the function may return 0, 2, 4, 6 or 7, as explained above.

Write an efficient algorithm for the following assumptions:
 N is an integer within the range [0..100,000];
 each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].


# Solution
The solution is discussed and provided in the resource.

Reendineering Python O(N) solution.
```python
def goldenLeader(A):
	n = len(A)
	size = 0
	
	for k in xrange(n):	
		if (size == 0):
			size += 1
			value = A[k]
		else:
			if (value != A[k]):
				size -= 1
			else:
				size += 1

	candidate = -1

	if (size > 0):
		candidate = value

	leader = -1
	count = 0

	for k in xrange(n):
		if (A[k] == candidate):
			count += 1
	
	if (count > n // 2):
		leader = candidate

	return leader
```
