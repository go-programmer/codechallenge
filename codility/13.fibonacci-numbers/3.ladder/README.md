# Problem

**Count the number of different ways of climbing to the top of a ladder.**

You have to climb up a ladder. 
The ladder has exactly N rungs, numbered from 1 to N. 
With each step, you can ascend by one or two rungs.

More precisely:
with your first step you can stand on rung 1 or 2,
if you are on rung K, you can move to rungs K + 1 or K + 2,
finally you have to stand on rung N.


Your task is to 
**count the number of different ways of climbing to the top of the ladder**.

For example, given N = 4, you have five different ways of climbing, ascending by:
1, 1, 1 and 1 rung,
1, 1 and 2 rungs,
1, 2 and 1 rung,
2, 1 and 1 rungs, and
2 and 2 rungs.

Given N = 5, you have eight different ways of climbing, ascending by:
1, 1, 1, 1 and 1 rung,
1, 1, 1 and 2 rungs,
1, 1, 2 and 1 rung,
1, 2, 1 and 1 rung,
1, 2 and 2 rungs,
2, 1, 1 and 1 rungs,
2, 1 and 2 rungs, and
2, 2 and 1 rung.

The number of different ways can be very large, 
so it is sufficient to 
**return the result modulo 2P, for a given integer P**.

Write a function:    func Solution(A []int, B []int) []int

that, given two non-empty arrays A and B of L integers,
returns an array consisting of L integers specifying the
consecutive answers; 

**position I should contain the number of different ways of climbing the ladder with A[I] rungs modulo 2B[I].**

For example, given L = 5 and:
A[0] = 4   B[0] = 3
A[1] = 4   B[1] = 2
A[2] = 5   B[2] = 4
A[3] = 5   B[3] = 3
A[4] = 1   B[4] = 1

4 mod 2*3 = 4 mod 6 = 2

the function should return the sequence [5, 1, 8, 0, 1], as explained above.

Write an efficient algorithm for the following assumptions:
L is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..L];
each element of array B is an integer within the range [1..30].


# Solution
At the first look the problem seems to be a
permutation problem with some modular arithmetic
operations.

It does not make a complete sense.

After searching for some answers, found a good
discussion at https://codesays.com/2014/solution-to-ladder-by-codility

Java solution was provided in the link. 
But still not able to understand the problem 
going through the solution.

```java
public int[] solution(int[] A, int[] B) {
        int[] fib = new int[A.length];
        int a = 0;
        int b = 1;
        int maxModulo = (int) Math.pow(2, 30);

        for (int i = 0; i < A.length; i++) {
            int x = (a + b) % maxModulo;
            fib[i] = x;
            a = b;
            b = x;
        }
        
        int[] modulo = new int[B.length];        
        for (int i = 0; i < B.length; i++) {
            modulo[i] = (int) Math.pow(2, B[i]);
        }
        
        int[] result = new int[A.length];
        for (int i = 0; i < A.length; i++) {
            result[i] = fib[A[i] - 1] % modulo[i];
        }
        
        return result;
    }
// https://app.codility.com/demo/results/demoR9XCB6-BQZ
```


Found another solution which is a bit more simple.
This solution also scored 100.
```js
function solution(A, B) {
    var i = 0;
    var max = 0;
    var maxB = 0;
    for(i=0; i<A.length; i++) {
        max = Math.max(max, A[i]);
        maxB = Math.max(maxB, B[i]);
    }
    

    var steps = [];
    steps[0] = 1;
    steps[1] = 1;
    i = 1;
    while(i++ <= max) {
        steps[i] = (steps[i-1] + steps[i-2]) % Math.pow(2, maxB);
    }
    
    var result = [];
    for(i=0; i<A.length; i++) {
        var div = steps[A[i]] & (Math.pow(2, B[i]) - 1);

        // Or
        // var div = steps[A[i]] % (Math.pow(2, B[i]));

        result.push(div);
    }
    
    return result;
}

// https://app.codility.com/demo/results/training65X9PC-YUB
// https://bloginstall.com/solution-to-codility-lesson-ladder-and-steps-or-rungs-javascript
```

In this solution the weird part is the `&` in the last for loop.
I've never used it so some details about it.

### Description 
&   Bitwise AND  
Returns 1 if the corresponding bits of both operands are 1, 
else it returns 0.

Performs binary and operation on operands binary value.

### Example 
12 & 25

12 = 01100
25 = 01001

00001100
00011001
---------
00001000  = 8 (In decimal)

https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_AND


a = 11;       
b = 10;       
console.log(a & b); 
console.log((a & b).toString(2)); //binary
//binary to decimal parseInt(binary, 2)

Checked with a slight change to use `%`. This
changed solution also scored 100. But it is
recommended to use & for large numbers.
```js
var div = steps[A[i]] % (Math.pow(2, B[i]));
// https://app.codility.com/demo/results/trainingMHJS6B-29Z
```


### Bitwise Operators
Operator 	Name 	                Description
& 	        AND 	                Sets each bit to 1 if both bits are 1
| 	        OR 	                    Sets each bit to 1 if one of two bits is 1
^ 	        XOR 	                Sets each bit to 1 if only one of two bits is 1
~ 	        NOT 	                Inverts all the bits

<< 	        Zero fill left shift 	Shifts left by pushing zeros in from the right,
                                    and let the leftmost bits fall off

>> 	        Signed right shift 	    Shifts right by pushing COPIES OF THE LEFT MOST BIT in from the left,
                                    and let the rightmost bits fall off

>>> 	    Zero fill right shift 	Shifts right by pushing zeros in from the left,
                                    and let the rightmost bits fall off