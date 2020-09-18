# Problem

Write a function:

    func Solution(A int, B int, K int) int

that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

    { i : A ≤ i ≤ B, i mod K = 0 }

For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

Write an efficient algorithm for the following assumptions:

        A and B are integers within the range [0..2,000,000,000];
        K is an integer within the range [1..2,000,000,000];
        A ≤ B.



# Analysis
One way to solve would be count the next number of lenght K.
For give example, 6..11, add 2 to 6 and again to the sum 8 
and so on. Here 6 is divisible by to so it seems straight 
forward.

Lets say K = 4, we will have to check if 6 is divisible by 4
as the count is inclusive of the end points.

Need to find if there is more efficient way to do it, probably
a mathematical solution.

Also we want to acheive the complexity below.
Complexity:
Expected worst-case time complexity is O(1);
Expected worst-case space complexity is O(1).


Found a mathematical solution:

Substraction of largest multiples of K.
B/K is the largest multiple of K 
A/K is the largest multiple of K


Substract the dividents from the larger number
and smaller number divided by divisor, K.
i.e. B/K - A/K

When the A/K is even, a divident will be excluded.
We need to add it.

( B/K - A/K ) + 1 (conditional)


### Resources:
https://www.youtube.com/watch?v=3985ygTPPq0
https://www.mathgoodies.com/lessons/vol3/divisibility
https://www.geeksforgeeks.org/count-numbers-divisible-m-given-range/
