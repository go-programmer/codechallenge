
# Problem

Write a function: func Solution(A []int) int
that, given a non-empty array A of N integers, 
returns **the minimal difference** that can be achieved
The difference between the two parts.

Write an efficient algorithm for the following assumptions:
- N is an integer within the range [2..100,000].
- Each element of array A is an integer within the range [−1,000..1,000].


For example, consider array A such that:
A[0] = 3
A[1] = 1
A[2] = 2
A[3] = 4
A[4] = 3


We can split this tape in four places:
P = 1, difference = |3 − 10| = 7
P = 2, difference = |4 − 9| = 5
P = 3, difference = |6 − 7| = 1
P = 4, difference = |10 − 3| = 7

The function should return 1, as explained above.


# Solution

### Analysis

The minimum sum will be 0,1 and so on..

##### Case 1: With negative number
Method 1: Direct sum
A[0] = -3
A[1] = 1
A[2] = 2
State Machine
1: -3 - (1+2) = -3 -3 = 6
2: (-3+1) - 2 = -2 -2 = 4


Method 2: Sort(A)
The array is already sorted.

Lets say the A was unsorted:
A[0] = 2
A[1] = 1
A[2] = -3
State Machine
1: 2 - 1-3 = 2 +2 = 4
2: 2+1 - -3 = 3 + 3 = 6

Sorted or unsorted, the sum will always be the same.

##### Case 2: Original solution with Sum
A[0] = 3
A[1] = 1
A[2] = 2
A[3] = 4
A[4] = 3

sum=13

sum - N
A[0] = 3 = 10
A[1] = 1 = 12
A[2] = 2 = 11
A[3] = 4 = 9
A[4] = 3


##### Case 3: Top-down approach on sorted data
A[0] = 3
A[1] = 1
A[2] = 2
A[3] = 4
A[4] = 3

Sort:
A[0] = 1
A[1] = 2
A[2] = 3
A[3] = 3
A[4] = 4


Sum = 13
State Machine: Top-down substraction.
P = 1(start), difference = |1 − (13-1)| = 12 - 1 = 11
P = 2, difference = |1+2 − 11| = 8 
P = 3, difference = |3 − 8| = 5
P = 4, difference = |4 − 5| = 1

For this set of data the solution is found at the end.

Add another element A[5] = 7.
Sum = 13 +7 = 20
P = 1(start), difference = |1 − (20-1)| = 19 - 1 = 18
P = 2, difference = |1+2 − 18| = 18 - 3 = 15 
P = 3, difference = |3+3 − 15| = 15 - 6 = 9
P = 4, difference = |6+4 − 9| = 10 - 9 = 1
P = 5, difference = |10+7 − 1| = 17 - 1 = 16

Here minimum is 1 at P=4. At this point sum is less than
aggregated sum. So the difference probably here will be 
the minimum, and should be returned.

##### Case 4: Bottom-up approach on sorted data
A[0] = 3
A[1] = 1
A[2] = 2
A[3] = 4
A[4] = 3

Sort:
A[0] = 1
A[1] = 2
A[2] = 3
A[3] = 3
A[4] = 4

Sum = 13
Sum = 13 - 4 = 9; Substract 4 from total or 
dont add the last number 4 in the total.

State Machine: Bottom-up substraction (or from end).
P = 0, difference = 1 - 1 = 0
P = 1, difference = 1 - 2 = 1
P = 2, difference = 2 - 3 = |-1| = 1
P = 3, difference = 5 - 3 = 2
P = 4(start), difference = |9 − 4| = 5

This process gives the minimum value of 0.
To note, at P=2, first number was lower than
second, 2,3. So could it be that the minimum
value might be it at this point when bottom-up
approach is taken?



##### Case 5: Thirdparty solution analysis
Complexity:
Expected worst-case time complexity is O(N);
Expected worst-case space complexity is O(N),
 beyond input storage (not counting the storage required for input arguments).

```golang
    arraySum := 0

    // Original solution code commented
    // and updated to make more easier sense.
    // And math package is imported for use.
    // currentMin := 1 << 32 - 1
	currentMin := math.MaxInt64

	for _, value := range A {
		arraySum += value
    }
    
    lhs := A[0]

    // Remove first value
	rhs := arraySum - lhs

	for i := 1; i < len(A); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))
        
        // Improvement
        // There is no need to check if this is true. 
        if diff == 0 return 0
        
		if diff < currentMin {
			currentMin = diff
        }
        
        // Increment sum from of left side.
        lhs += A[i]
        
        // Decrease current value of left side.
		rhs -= A[i]
	}

	return currentMin
```

This solution get's a 100% score.

The other alternatives with sort could be effective
and needs an implementation to test. 

For now I will just put it to a TODO list.
