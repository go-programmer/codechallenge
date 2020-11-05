# Reformat digits into blocks of lenght 3 separated by dash

00-44  48 5555 8361 = 004-448-555-583-61
00-44  48 5555 8361 = 004-448-555-583-61

0 - 22 1985--324 = 022-198-53-24
555372654 = 555-372-654

We are given a string S representing a phone number, 
which we would like to reformat. 

String S consists of N characters: digits, spaces and/or dashes. 
It contains at least two digits.

Spaces and dashes in string S can be ignored.
We want to reformat the given phone number in
such a way that the digits are grouped in blocks
of length three, separated by single dashes.

# If necessary, the final block or the last two blocks can be of length two.
`Need to confirm this, and appropriate example is not given`


Focus on correctness.

# Solution
Make a group of numbers of lenght 3.
Read every character.

If it is a space or dash, then ignore.

If it is a number then append to the string.

We need to add a dash after the number if the lenght is two.









