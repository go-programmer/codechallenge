# Problem

A non-empty array A consisting of N integers is given.

**A peak is an array element which is larger than its neighbours.**

More precisely, it is an index P such that 
**0 < P < N − 1** and 

**A[P − 1] < A[P] > A[P + 1]**.

For example, the following array A:
    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2

has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains 
whose relative heights are represented by array A,
as shown in a figure below. 

**You have to choose how many flags you should take with you.** 

The goal is to **set the maximum number of flags on the peaks**, according to certain rules.

https://codility-frontend-prod.s3.amazonaws.com/media/task_static/flags/static/images/auto/6f5e8faa3000c0a74157e6e0bc759b8a.png



**Flags can only be set on peaks**.
What's more, if you take K flags, 
then **the distance between any two flags should be greater than or equal to K**.

The distance between indices P and Q is the absolute value |P − Q|.

For example, given the mountain range represented by array A, 
above, with N = 12, if you take:
  two flags, you can set them on peaks 1 and 5;
  three flags, you can set them on peaks 1, 5 and 10;
  four flags, you can set only three flags, on peaks 1, 5 and 10.

You can therefore set a maximum of **three flags** in this case.

Write a function: func Solution(A []int) int

that, given a non-empty array A of N integers, 
returns the maximum number of flags that can be set on the peaks of the array.

For example, the following array A:
    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2

the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [1..400,000];
        each element of array A is an integer within the range [0..1,000,000,000].


# Solution
The goal is to **set the maximum number of flags on the peaks**


### Find peaks 
A peak is an array element which is larger than its neighbours: A[P − 1] < A[P] > A[P + 1]
and 0 < P < N − 1


### Set flag.

The first and last wont be peaks and wont have flag set.

A has exactly 4 peaks: elements 1, 3, 5 and 10.
A[0] = 1
A[1] = 5 P F
A[2] = 3
A[3] = 4 P 
A[4] = 3
A[5] = 4 P F
A[6] = 1
A[7] = 2
A[8] = 3
A[9] = 4
A[10] = 6 P F
A[11] = 2


Then K flags on the peaks. The set flag is 3.
There are conditions to set flag:
  **The distance between any two flags should be greater than or equal to K.**
  The distance between indices P and Q is the absolute value |P − Q|.

### Observation
The maximum number of flags will be equal to or less than the number of peaks.


We need to find the maximum distance between the peaks.
A[1]  = P F
A[3]  = P 
A[5]  = P F
A[10] = P F

Minimum distance = 1-3 = 2
Maximum distance = 1-10 = 9

Lets say A array of distance among the indices.
A = {1,3,5,10}
What we want to do is break A down into as many
sub-arrays as possible with max length of those 
sub-arrays less than or equal to 4. 

# Theory
Could it be we can take the factorial as there are 
4 peaks and the factorial of are the number of flags?
3, i.e. 1,2, and 4.
This theory failed miserably with 26% score.

#### Find possible flags that can be set.
Lets try to set flags in all of the peaks, 4.
Total distance to set 4 flags should be at least
4*4 = 16, to satisfy the first condition to set
flag. The total distance here is 9. So we move
on to set 3 flags. 

To set 3 flags, the total distance needed is 9,
which is provided. Lets proceed forward to check 
if there are values that satifies the conditions.

maxPossibleFlag = total peaks
if total distance >= (total peaks * total peaks) 
        maxPossibleFlag = total peaks

#### Set flag
P[0] = 1 F
P[1] = 3  
P[2] = 5 F
p[3] = 10 F

flag=1,
flag is set to 1, why? 

If there is at least one peak,
there will be one flag possible.

P[0]=1
P[1]=3 One peak, one flag
P[2]=2


for i=1; i < maxPossibleFlag; i++ 

        if P[i+1] - A[i] <= maxPossibleFlag 
                flag++

        else if A[i+2] - A[i] <=3
                flag++
                i++  

### Change mindset
"panic: runtime error: index out of range"
solution.countFlags(0xc208066000, 0xd05, 0x2710, 0xc208066000)
/tmp/workspace/src/solution/solution.go:56 +0xad

is error thrown at the last attempt.
https://app.codility.com/demo/results/trainingSSBQM4-GFS/

The condition might be causing the issue:

```go
if peaks[i+1]-peaks[i] >= peakLen {
        flag++

} else if peaks[i+2]-peaks[i] >= peakLen {
        flag++
        i++
}
```

Need to think about fiding other ways of finding max peaks.


### Conclusion
Tried out many variations but the final solution with max result could only be
reached with the logic for fiding max flags counts using the thirdparty solution.

The max score reached ws 86 percent.

