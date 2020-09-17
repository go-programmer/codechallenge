# Problem

You are given N counters, initially set to 0, and you have two possible operations on them:
increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.

A non-empty array A of M integers is given. This array represents consecutive operations:
if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.

Write a function:
    func Solution(N int, A []int) []int
that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

Result array should be returned as an array of integers.

For example, given integer N = 5 and array A such that:
    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4

the values of the counters after each consecutive operation will be:
    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)

the function should return [3, 2, 2, 4, 2], as explained above.

# Analysis
I could not understand the problem at first.

Reviewed a solution that socred 100, still could 
understand the problem.

To understand the problem, watched video at:
https://www.youtube.com/watch?v=H47iCG2YiA0&list=TLPQMTcwOTIwMjCxoGhdzxe17w


So intstruction/counter table can be formed as:

Instruction     Counter N = 5    
    A           1  2  3  4  5
    3           0  0  1  0  0
    4           0  0  1  1  0
    4           0  0  1  2  0
    6           2  2  2  2  2
    1           3  2  2  2  2 
    4           3  2  2  3  2
    4           3  2  2  4  2


# Solution

Create an array of size N, 
counter [6]int

Iterate A:

Check conditions:

If A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X):
Increment counter[A.value]++

If A[K] = N + 1 then operation K is max counter.
If A.value > N, set counter[1 to 5] to A.value.
This step will take quadraric time if we do it 
with a simple inner loop.


Lets say we note the max, set it to counter[0].

Instruction     Counter N = 5    
    A           1  2  3  4  5
    3           0  0  1  0  0
    4           0  0  1  1  0
    4           0  0  1  2  0
    6           2  2  2  2  2 counter[0] = 2

For next instruction, we need to check
if the value of counter[0] > counter[1].
If true, set the value of it by increasing
the counter[0] value.

    1           3  2  2  2  2 
    
From this condition will have to be checked
for all instructions.

    4           3  2  2  3  2
    4           3  2  2  4  2

Suppose another instruction is 7.
In the last instruction, we the value of
counter 4 was 4, so we instead of going 
through the array to find the max value
we should have checked if any of the counter 
value has increased more than counter[0].

In any increment operation check if max value 
is greater than counter[0]?

counter[A.value]++
if counter[A.value] > counter[0]
    maxCounter = counter[A.value]

Now, if later there is a max counter operation then
Check maxCounter > counter[0],
    counter[0] = maxCounter


The first solution scored 44%.