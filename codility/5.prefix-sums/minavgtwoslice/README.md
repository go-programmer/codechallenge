# Problem



A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains **at least two elements**). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8

contains the following example slices:

        **slice (1, 2), whose average is (2 + 2) / 2 = 2;**
        slice (3, 4), whose average is (5 + 1) / 2 = 3;
        slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.

The goal is to **find the starting position** of a slice whose average is minimal.

Write a function:

    class Solution { public int solution(int[] A); }

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is **more than one slice with a minimal average**, you should return **the smallest starting position of such a slice**.

For example, given array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1    
    A[5] = 5
    A[6] = 8

the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [2..100,000];
        each element of array A is an integer within the range [−10,000..10,000].


# Analysis

State Machine

            P
Array       PrefixSum  Avg.From0  Avg.From1        Avg.From2       Avg.From3   Avg.From4   Avg.From5
A[0] = 4    4           - skip 1  -                 -              -           -           -

A[1] = 2    6           6/2       - skip 1          -              -           -           -
                        =3

A[2] = 2    8           8/3        8-4|2+2/2        - skip 1       -           -           -
                        =2.6       =**2** 

A[3] = 5    13          13/4       13-4|2+2+5/3     13-6|2+5/2      skip 1     -           -
                        =3.25      =9/3=3           =7/2=3.5 

A[4] = 1    14          14/5       14-4|2+2+5+1/4   14-6|2+5+1/3    14-8/2      skip 1     -
                        =2.8       10/4=2.5         8/3=2.67        6/2=3

A[5] = 5    19          19/6       19-4|2+2+5+1+5/5 19-6/4          19-8/3      19-13/2     skip 1
                        =3.17      =15/5=3          =13/4=3.25      =11/3=3.67  =6/2=3      

A[6] = 8    27          27/7       27-4/6           27-6/5          27-8/4      27-13/3     27-14/2
                        =3.86      =23/6=3.83       =21/5=4.2       =19/4=4.75  =14/3=4.67  =13/2=6.5



1. Take prefix sum.

2. Get minimum average.
We cannot use Greedy approach and skip remaining
as there is the requirement to take minimum point 
in case of two min avg is found, which is a possibility.
Process for this:
This will be quadratic as we have to loop through average 
values from every index.

Lets say min average is set to Int.Max
Loop in the prefix sum array.
Another loop to find the average.

Completed the above state machine for further analysis.

First loop, 
- We skip P[0], start from P[1]
- Divide P[1]/2, P[2]/3 ... P[len-1]/len-1

Second loop, 
- We skip P[0], P[1], i.e start from P[2]
- Divide P[2]/2, P[3]/3 ... P[len-1]/len-1


### Pseudocode 
- Loop P until len-1,
- Inner loop starts from next step
    - Skip previous step
    - Division starts by dividing by 2, and increases

Is there any N or logN approach?

The initial solution scored 0, and failed with time out error.
Result: https://app.codility.com/demo/results/trainingSEXCGZ-3XE/

After some research, found out that this is a mathematical problem.
Also the problem did not state time complexity, i.e.

Complexity:
    expected worst-case time complexity is O(N);
    expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).

If the time complexity had been given I 
would have thought of mathematical answer.

The assumptions in the problem also did not
provide hints as it seemed quadratic time 
solution was acceptable, i.e.

Assumption:
    N is an integer within the range [2..100,000];
    each element of array A is an integer within the range [−10,000..10,000].


Anyways, from the analysis of the state machine,
it was somewhat clear that more the steps proceeded,
the minimum average increased. So the solution had 
to come with in first few steps.

# I overthought this problem 


`Each element of array A is an integer within the range [−10,000..10,000].`
This assuption pushed me to get all the sums.

Tried a thirdparty soluttion which passed the tests. 
And the comment pushed me to search more mathematical 
definitions.

Result: https://app.codility.com/demo/results/trainingVEKHW6-32B

// There is a mathematical rule that says you only have to find the min of 2 or 3 consecutive slices
```golang
func MinAvgTwoSlice(A []int) int {
	arrayLen := len(A)
	minAvg := float64(1<<32 - 1)
	result := 0

	for i := 0; i < arrayLen-1; i += 1 {
		if i+1 < arrayLen {
			temp := (float64(A[i]) + float64(A[i+1])) / float64(2.0)

			if temp < float64(minAvg) {
				minAvg = temp
				result = i
			}
		}

		if i+2 < arrayLen {
			temp := (float64(A[i]) + float64(A[i+1]) + float64(A[i+2])) / float64(3.0)

			if temp < float64(minAvg) {
				minAvg = temp
				result = i
			}
		}
	}

	return result
}
```

Found a good solution description here:
https://medium.com/@molchevsky/best-solutions-for-codility-lessons-lesson-5-prefix-sums-68b716f9d825

It provides good insights.


Found a solution with using prefix sum and similar concept.
```java
class Solution {
    public int solution(int[] A) {
        int result = 0;
        int N = A.length;
        int [] prefix = new int [N+1];
        for (int i = 1; i < prefix.length; i++) {
            prefix[i] = prefix[i-1] + A[i-1];
        }
        double avg = Double.MAX_VALUE;
        for (int i = 1; i < N; i++) {
            for (int j = i+1; j <=N; j++) {
                double temp = (double)(prefix[j]-prefix[i-1]) /(double)(j-i+1);
                if (temp < avg) {
                    avg = temp;
                    result = i-1;
                }
            }
        }
        return result;
    }
}
```
It scored 60% while running. For result:
Detected time complexity:O(N ** 2) 
https://app.codility.com/demo/results/trainingNCBH8B-4FT

The original result was also had also scored 60% but 
differed in the Detected time complexity: O(N ** 3) or O(N ** 2)
https://www.xspdf.com/help/50605322.html 
https://app.codility.com/demo/results/demo65RNV5-T36


Found another solution https://funnelgarden.com/minavgtwoslice-codility-solution
which scored 100%. The concept is similar though.
Result: https://app.codility.com/demo/results/trainingG4KGVB-23J
```java
public class MinAverageTwoSlice {
  public int solution(int[] A) {

    //main idea: will find min average by checking only 2 and 3 contiguous elements at a time
    int sum1, sum2 = 0;
    double minAverage = Double.MAX_VALUE;
    double currentAverage1 = Double.MAX_VALUE;
    double currentAverage2 = Double.MAX_VALUE;
    int minAverageSliceIndex = 0; //for size 2 arrays, this will always be true

    //if array is > 2 elements
    for(int i=0; i<A.length-2; i++) {
      sum1 = A[i] + A[i+1];
      currentAverage1 = sum1 / 2.0d;
      if(currentAverage1 < minAverage) {
        minAverage = currentAverage1;
        minAverageSliceIndex = i;
      }

      sum2 = sum1 + A[i+2];
      currentAverage2 = sum2 / 3.0d;
      if(currentAverage2 < minAverage) {
        minAverage = currentAverage2;
        minAverageSliceIndex = i;
      }
    }

    //check last 2 contiguous elements from the end - they won't otherwise be checked because
    //when checking 2 and 3 contiguous elements at a time, will stop checking 3 elements from the end
    currentAverage1 = (A[A.length-2] + A[A.length-1]) / 2.0d;
    if(currentAverage1 < minAverage) {
      minAverage = currentAverage1;
      minAverageSliceIndex = A.length-2;
    }
   
    return minAverageSliceIndex;
  }
}
```


Well the usage of 2 and 3 as divisor is intersting.
More research required.