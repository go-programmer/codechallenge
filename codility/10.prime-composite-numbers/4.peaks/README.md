# Problem

Divide an array into the maximum number of same-sized blocks,
each of which should contain an index P such that 

A[P - 1] < A[P] > A[P + 1]. 


A non-empty array A consisting of N integers is given.

**A peak is an array element which is larger than its neighbors.**

More precisely, it is an index P such that 0 < P < N − 1,  
A[P − 1] < A[P] and A[P] > A[P + 1].

For example, the following array A:
    A[0] = 1
    A[1] = 2
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

has exactly three peaks: 3, 5, 10.

**We want to divide this array into blocks containing the same number of elements.** 

More precisely, we want to choose a number K that will yield the following blocks:

        A[0], A[1], ..., A[K − 1],
        A[K], A[K + 1], ..., A[2K − 1],
        ...
        A[N − K], A[N − K + 1], ..., A[N − 1].

What's more, **every block should contain at least one peak.**

Notice that extreme elements of the blocks (for example A[K − 1] or A[K]) can also be peaks,
but only if they have both neighbors (including one in an adjacent blocks).

**The goal is to find the maximum number of blocks into which the array A can be divided.**

Array A can be divided into blocks as follows:

I: 0 1 2 3 4 5 6 7 8 9 10 11
V: 1 2 3 4 3 4 1 2 3 4 6  2

One block (1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2). 
This block contains three peaks.

Two blocks (1, 2, 3, 4, 3, 4) and (1, 2, 3, 4, 6, 2). 
Every block has a peak.

Three blocks (1, 2, 3, 4), (3, 4, 1, 2), (3, 4, 6, 2). 
Every block has a peak. Notice in particular that 
the first block (1, 2, 3, 4) has a peak at A[3], 
because A[2] < A[3] > A[4], even though A[4] is in 
the adjacent block.

However, array A cannot be divided into four blocks, 
(1, 2, 3), (4, 3, 4), (1, 2, 3) and (4, 6, 2), 
because the (1, 2, 3) blocks do not contain a peak.

Notice in particular that the (4, 3, 4) block contains
two peaks: A[3] and A[5].

The maximum number of blocks that array A can be 
divided into is three.


Write a function: func Solution(A []int) int

that, given a non-empty array A consisting of N integers,
returns the maximum number of blocks into which A can be divided.

If A cannot be divided into some number of blocks, the function should return 0.

For example, given:
    A[0] = 1
    A[1] = 2
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
N is an integer within the range [1..100,000];
each element of array A is an integer within the range [0..1,000,000,000].


# Solution

The first thing to determine is peak, which is 
A[P - 1] < A[P] > A[P + 1], same as 5.

The main part of the problem is to determine the
maximum number of same-sized blocks where each 
block should contain a peak.

The maximum number of the blocks would be the 
number of peaks or less than peaks, i.e.
blocks <= peaks 

Lets first check if A can be divided into max
blocks. What this also means is that the distance
between each peak must be equal.


I: 0 1 2 3 4 5 6 7 8 9 10 11
V: 1 2 3 4 3 4 1 2 3 4 6  2
P: 0 0 0 1 0 2 0 0 0 0 3  0
D:         2           5  

First we need to think about would be possible that
if the A can divided into maximum, i.e. blocks = peaks?

The naive way to check would be iterate through A
and check for every block if there exists a peak.

If peaks is not possible then reduce peak by 1 
and check the same process as before.

Lets take distance concept.
To divide A into max peak, length of A must be
greater than square of max peak. In the provided
case,
peak = 4
length A = 12

peak * peak > len A, so 4 is not possible.


Now, reduce peak, peak=3, peak*peak=9, which is 
less than 12, so 3 blocks are possible.

We need to determine the size of the block, that 
would be lenght of A divided by peak, i.e.
size=lenA/peak, here size=4. 

We need to find if each block size has at least one peak.

```golang

blockSize = 4
for i:=0; i<len(A); i++ {
	
	peakExistsCounter = blockSize

	Check if peak exists in blockSize
	for j = peakExistsCounter;j>-1;j++ {
		
		We need to store all the peaks and their index		
		
		if peak exists in this block then,
			Reduce total peaks -- 

            // OLD CONCEPT
			Reset the i to starting point of next block.
			Say peak is found at 0, next block would be
			i+4-0. If peak at 1, then next block would be
			i+4-1. In general, i+4-j.
			
			Lets suppose, the peak does not exits in this size,
			but size+1 or two.

			And also the lenA is 15 where the block size is still 4
			and flag count 3 as provided.

            The above logic will break.
            
            The blocks must be contigious not to break the flow. 
            So when one peak is found, we need to take include the
            maximum previous index in order to create the block of
            that size. 

            //NEW THOUGHT
            Block sizes will always be fixed. We just need to check
            if a peak exists in the block.

    
Index: 0 1 2 3 4 5 6 7 8 9 10 11

Value: 1 2 3 3 4 3 1 4 3 4 6  2

Peaks: 0 0 0 0 1 0 0 2 0 0 3  0
            
Here first peak is at 4ht index. After we
find the first block, we find second peak
and check if it is in block of size required.

But in above A, the last peak does not have the 
size required as it has only 3 in the block.

So we need to check for 2 blocks.


    }
	  

}
```

### Another Perspective
For A of len L, the maximum block size that can be divided
to its factorial x*y. And check if x blocks of size y, is possible
or not and y block of x. and other factorial.

For A len 15, the maximum block size that can be divided is its factorial
3*5=15, 5 blocks of size 3, and 3 blocks of size 5.

Given, For A len 12, 3*4, 3 blocks of size 4

121			L3 P1 Prime/Odd
1212		L4 P1
12121  		L5 P2 Prime/Odd	+2
121212 		L6 P2
1212121		L7 P3 Prime +2
12121212	L8 P3			L is even, so the max number of possible P is
121212121 	L9 P4 Not Prime/Odd +2

1212121212 		L10 P4
12121212121		L11 P5 Prime/Odd +2
121212121212	L12 P5
1212121212121	L13 P5 Odd/ +2

From the tren, there is an peak or probality of peak in every odd number.

But this trend will not always happen as it will be more random but
the number of box possible will be in terms of factorial.

**The minimum size of the block size could be two.**

### Checking if a block contains a peak
I: 0 1 2 3 4 5 6 7 8 9 10 11
V: 1 2 3 4 3 4 1 2 3 4 6  2
P: 0 0 0 1 0 2 0 0 0 0 3  0

lenA = 12
peaks = 3

Maximum Number of blocks = 12/3 = 4. But there are only 3 peaks.
So we need to check if 3 blocks are possible which will have the
size 4.

Suppose there are 4 peaks then we also have to check for 4 block
of size 4. So there will be two loops for a factor of the number.


here, block = 4

for lenPeaks*lenPeaks < lenA {

    // To divide A into equal number of blocks with at leat one
    // peak in the block, lenPeak must be a factor of lenA.
    if lenA%lenPeaks == 0 {

        // lets do this in second refactoring
        // first lets check for larger block
        // lenA%block 

        b - 0-3,4-7,8-11
        p - 3,5,10

        block =  lenA/lenPeaks 12/3=4 
        hasPeak == false
        peakIndex=0

        As we have stored the indexes of peaks, we can use it check if a peak exists in the block.
        j=0; 
        for j < lenA 12; {

            min = j, max = j+4

            // how many peaks falls under min and max
            // at least one is needed.
            for k=min; k < max; k++{

                if peaks[peakIndex] == i {
                    peakIndex++
                    hasPeak = true
                }
            }

            // if no peak exists in the block
            if hasPeak == false {
                break 
            }
            
            // if all peaks are finished but more blocks
            // are left.
            if peakIndex > lenPeaks && j < lenA {
                break
            }

            // Check next block
            j + block 4 
        }

        
        // How to know if all the blocks has a peak?
        // We have looped into every block and checked if the
        // peak exists in each.

        // All blocks contains at least one peak
        if peakIndex == lenPeaks && j == lenA {
            return blocks
        }
    }

    peak--
}



We need tget the factorial to get the maximum blocks
size we can create compared to the number of peaks we
have.

Also having the the minimum distance between two peaks,
we know how small size block we can create.
For eg, 121212, here we can create a 2 block of size 3


The first solution scored 27 and failed
For example, for the input [0, 1, 0, 0, 1, 0, 0, 1, 0]
the solution returned a wrong answer (got 0 expected 3). 
https://app.codility.com/demo/results/trainingENW8Y5-JNS

The second solution scored 36
https://app.codility.com/demo/results/trainingU4TP2J-VZQ


Third solution scored 54:  
Task Score 54%, Correctness 100%, Performance 0
https://app.codility.com/demo/results/trainingMYXRMK-YM5

The last solution retured 1 for all cases:
got 1 expected 100
got 1 expected 625 
got 1 expected 5000
got 1 expected 30030
got 1 expected 25000


There was an error in the countPeaks() while 
checking for odd number. Removed that logic
and checked for other cases but the score 
remained same.
https://app.codility.com/demo/results/training3MR87C-NM4


### Different approach.
The number of ways a number can broken down with equal parts
is by its factorial like 6=2*3, either 2 blocks with 3 elements
or 3 blocks with 2 elements. The size of block is determined by
number of peaks.

To have maximum number of blocks, i.e. 3, there needs to be at
least 3 peaks in those two blocks.

We can get the factorials and start from the range according to
the number of peaks


```golang
// Returns highest divisor of n
// O(√n)
func highestDivisor(n int) int {
	i := 1
	highestDivisor := 0
	divisors := make([]int, n/2)
	divisorsLen := 0

	for i*i < n {

		if n%i == 0 {
			highestDivisor = i
			divisors[divisorsLen] = i
			divisorsLen++
		}

		i++
    }
    
    divisors = divisors[:divisorsLen]
    fmt.Println(divisors)
	return highestDivisor
}
```

For A = 100000000, the factorials will be:
1 2 4 5 8 10 16 20 25 32 40 50 64 80 100 125
128 160 200 250 256 320 400 500 625 640 800
1000 1250 1280 1600 2000 2500 3125 3200 4000 
5000 6250 6400 8000

We can iterate through this range and check 
accordingly.