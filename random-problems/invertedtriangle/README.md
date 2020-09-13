
# Problem
Print stars in the inverted triangle format.

        * * * * * * *
          * * * * *
            * * *
              *

# Solution

### Analysis

> Pattern found
1. prints all n stars
2. prints n-1 stars, skip one in each end
3. prints n-2 stars, skip two in each end
...

Q. How to print until one star? Or how to  know when to stop.
Q. How tall will be the triangles height?


Q. Is this format possible for even numbers?
8 stars
* * * * * * * * 
  * * * * * *
    * * * *
      * *
       *

6 stars
* * * * * * 
  * * * *
    * * 
     *

4 stars
* * * *
  * * 
   *

Looks like it's not possible to make a good triangle
with even numbers.

Only with even numbers, an inverted triangle is possible.


When n=7; height=4, skipped start increases by 1.
* * * * * * *
  * * * * *
    * * *
      *

When n=5; height=3
* * * * *
  * * *
    *

When n=11; height=6, stars=n-2, space=( n - current n )/2
* * * * * * * * * * * 11
  * * * * * * * * * 9 stars, 2 space
    * * * * * * * 7 stars, 4 space
      * * * * * 5 stars, 6 space
        * * * 3 stars, 8 space
          * 1 star, 10 space

* * * * * * * 7 stars, 0 space
  * * * * * 5 stars, 2 space
    * * * 3 stars, 4 space
      * 1 star, 6 space

Hence, height = (n/2)+1
And the number of starts to print = (n - current n)/2 

> Pseudocode
height = (n/2)+1
Loop height > 0; height--
  starts=n
  Loop starts > 0; starts--
    print *
  println
  n-2
