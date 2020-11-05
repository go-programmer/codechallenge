
# Problem:
Determine whether a triangle can be built from a given set of edges. 


An array A consisting of N integers is given.

A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:
  A[P] + A[Q] > A[R],
  A[Q] + A[R] > A[P],
  A[R] + A[P] > A[Q].

For example, consider array A such that:
  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 20

Triplet (0, 2, 4) is triangular.

Write a function:

    func Solution(A []int) int

that, given an array A consisting of N integers, **returns 1 if there exists a triangular triplet** for this array and returns 0 otherwise.

For example, given array A such that:
  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 20

the function should return 1, as explained above. Given array A such that:
  A[0] = 10    A[1] = 50    A[2] = 5
  A[3] = 1

the function should return 0.

Write an efficient algorithm for the following assumptions:
    N is an integer within the range [0..100,000];
    each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].


# Solution

Conditions for triangular number to exist, after sort:
A[0] + A[1] > A[2]
A[1] + A[2] > A[0]
A[2] + A[0] > A[1]



The trivial solution would be quadratic:

Sort A

First loop:
Check A[0]+A[1] > A[2]
Keep A[0], increment A[index++]

Second loop:
Increment base, A[1]+A[2] > A[3]

Looking at different test cases:

Case: 1,2,3
1+2 > 3, No

Case: 1,2,4
1+2 > 4, No

Case: 2,3,4
2+3 > 4, Y; 3+4 > 2, Y; 4+2 > 3, Y

Case: 2,2,5
2+2 > 5, n

Case: 2,4,5
2+4 > 5, y
4+5 > 2, y
5+2 > 4, y


Example Case 2: 10, 50, 5, 1
1,5,10,50
1+5 > 10, n

move index
1+10 > 50, n

move index pointer

5+10 > 50, n; return 0


Random case: 1 	3 	6 	8
1+3>6,n

Difference in the numbers 
3-1=2, 6-3=3, 8-6=2

The diff is decreased at end.
3+6>8,y


Example Case 1: 10, 2, 5, 1, 8, 20
1 2 5 8 10 20

1+2 > 5, n ; Diff: 1+2=3, 2+5=7      | 2-1=1, 5-2=3
2+5 > 8, n ; Diff: 2+5=7, 5+8=13     | 5-2=3, 8-5=3 
5+8 > 10, y ; Diff: 5+8=13, 8+10 =18 | 8-5=3, 10-8=2 * Here the difference has decreased

At sub-sequence 5 8 10
Taking the difference between two next consecutive 
numbers 10-8 < 5. This could be a way to think.
8-5  = 3
10-8 = 2

At sub-sequence 5 8 10 11, for 11
8 - 5 = 3
12 - 8 = 4 

Need to check for negative mix.

Case: -1,-3,-5
-5 -3 -1 





Found similar concept implementation in Java
```java
// write your code in Java SE 8
int n = A.length;
if(n<3){
    return 0;
}
Arrays.sort(A);
for(int i=2; i<n; i++){
    if(A[i]<(long)A[i-1]+(long)A[i-2])
    return 1;
}
return 0;
```

# Conclusion
We just needed to check the first condition after sorting
the array as the rest of the conditions would automatically 
be true.

i.e. A[0] + A[1] > A[2]

