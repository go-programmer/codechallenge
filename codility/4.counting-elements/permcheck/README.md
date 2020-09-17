# Problem
A permutation is a sequence containing each 
element from 1 to N **once, and only once**.

For example, array A such that:
    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2

is a permutation, but array A such that:
    A[0] = 4
    A[1] = 1
    A[2] = 3

is not a permutation, because value 2 is missing.

Write a function:

    func Solution(A []int) int

that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

Assumptions:
    N is an integer within the range [1..100,000];
    Each element of array A is an integer within the range [1..1,000,000,000].


# Analysis
The number should be unique, no repetition.

The number should be continuous:
it must start from 1, and last number
must be lenght of A == A[len(A)].

We take a sort first apporach, then check.

Sort A

If A[0] != 1 || len(A) != A[len(A)]
    return 0

With checking len(A) we will know if the 
last number is exists and validate continuity
but that does not confirm all other numbers
in between exists. For example, A[1,2,3,5,5]
does not contain 4.

Now we need to check for the next number.
```golang

// Set a counter to compare with the number
// in the array.
count := 2
// Skip 1st and last
for a range A[ 1 : len(A)-1] {

    if a != count {
        return 0
    }

    count++
}

return 1
```

