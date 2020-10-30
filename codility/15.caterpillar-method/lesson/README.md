
https://app.codility.com/programmers/lessons/15-caterpillar_method

# Caterpillar Method

The idea is to check elements in a way that’s
reminiscent of movements of a caterpillar.
The caterpillar crawls through the array.
We remember the front and back positions 
of the caterpillar, and at every step either 
of them is moved forward.


15.2. Exercise
Problem: You are given n sticks (of lengths >= 1 and <= 10^9).
The goal is to count the number of triangles
that can be constructed using these sticks.

More precisely, count the number of triplets
at indices x < y < z, such that A[x] + A[y] > A[z].

Solution O(n^2):
For every pair x, y we can find the largest stick z that can be used to
construct the triangle. Every stick k, such that y < k <= z, can also be used, because the
condition a x + a y > a k will still be true. We can add up all these triangles at once.

If the value z is found every time from the beginning then we get a O(n^3) time complexity
solution. However, we can instead use the caterpillar method. When increasing the value of
y, we can increase (as far as possible) the value of z.


