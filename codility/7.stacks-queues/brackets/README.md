
# Problem:
Determine whether a given string of parentheses (multiple types) is properly nested.


A string S consisting of N characters is considered
to be properly nested if any of the following conditions is true:
S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.

For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function: func Solution(S string) int
that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [0..200,000];
        string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

# Solution
Read the string from the beginning
Push into string if "(", "{", "[".
Pop when string is "]", "}", ")". 
Check if the closing bracket matches the closing one.
When the loop ends the index of the array should be 0

Slice implementation of stack
String: {[()()]}
Slice :A[]
Loop

Read first character: {
Check for push, pop operation
Push into slice, mark slice index
Index++

Read char: [ 
Push/Pop Op: Push

Read char:( 
Push/Pop Op: Push


Read char: ) 
Push/Pop Op: Pop
Get last inserted value (
Open close brackets matches, continue.



Character values of the brackets
{ 123 125 }

[ 91 93 ]

( 40 41 )


The first solution scored 50%.
The solution failed for empty string.
There was runtime error: index out of range.
I ignored other string values.


Second solution also failed, scored 50%.

Third solution progressed with 62% score.
Passed empty string condition.
An example of failed case:
For example, for the input '())(()' the solution terminated unexpectedly. 

Fourth improved score to 75%.
Passed test  case multiple_full_binary_trees.
Failed example: for the input '{{{{' the solution terminated unexpectedly.

Fifth solution also scored 75% but Correctness got 100%

Sixth solution socred 87%. This solution like 7th failed 
a test case that previous ones passed:
large1 : simple large positive test, 100K ('s followed by 100K )'s + )(


7th solution with minor changes also scored same and same problem.

Reviewed previous solution which passed that problem, and made changes.

8th solution failed 
large2: simple large negative test, 10K+1 ('s followed by 10K )'s + )( + ()

Looks like when large 1 passes large 2 fails and vice versa.


9th solution scored 100%. Increased the stack size to len of array.

Optimized 9th solution as the memory usage was maximum. Reduced it to half.

Optimized more by making slice of stack of int8, from byte(uint8).
