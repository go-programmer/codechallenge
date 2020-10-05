# Probolem


Two positive integers N and M are given. 
Integer N represents the number of chocolates
arranged in a circle, numbered from 0 to N − 1.

You start to eat the chocolates. After eating 
a chocolate you leave only a wrapper.

**You begin with eating chocolate number 0**. 
Then you 
**omit the next M − 1 chocolates or wrappers on the circle, and eat the following one.**

More precisely, if you ate chocolate number X, 
then you will next eat the chocolate with number
**(X + M) modulo N (remainder of division)**.

You **stop eating when you encounter an empty wrapper**.

For example, given integers N = 10 and M = 4.
You will eat the following chocolates: 0, 4, 8, 2, 6.

The goal is to **count the number of chocolates that you will eat**,
following the above rules.


Write a function:    func Solution(N int, M int) int

that, given two positive integers N and M, returns the number of chocolates that you will eat.

For example, given integers N = 10 and M = 4. the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:
        N and M are integers within the range [1..1,000,000,000].

# Solution
N = 10 and M = 4; next choclate (X + M) modulo N or M − 1;
You will eat the following chocolates: 0, 4, 8, 2, 6.

State Machine
N   Next        Mod
0   4           0+4%10=4
4   8           4+4%10=8
8   (9,0,1),2     8+4%10=2
2   (3,4,5),6     2+4%10=6
6   (7,8,9),0     6+4%10=0

Next is 0 which is already eaten, we stop.

### Thought 1
Loop N
Find Next, 
    If N exists in map return count
    Else, store n in the map

This does not use the lesson concepts, GDC/LCM.

### Thought 2
The two numbers are 10 and 4
GCD 2

1: 10/2 = 5

LCM 10*4 / 2 =20
2: LCM/4 = 5


Simple division by GCD, i.e. 1, worked and 
the solution scored 100.

> More clarfication on why GCD usage worked here:
Further more, rearranging the choclates being eaten
in ascending order 0 2 4 6 8, the numbers increasing
order by 2. In a way the choclates eaten were at those
indexes.

