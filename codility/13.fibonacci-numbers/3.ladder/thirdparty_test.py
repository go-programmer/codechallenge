
# https://docs.python.org/3/library/unittest.html

import unittest

##class Solution:
def solution(A, B):
    limit    = max(A)                 # The possible largest N rungs
    result   = [0] * len(A)           # The result for each query
    modLimit = (1 << max(B)) - 1      # To avoid big interger in fibs
    # Compute the Fibonacci numbers for later use
    fib    = [0] * (limit+2)
    fib[1] = 1

    for i in range(2, limit + 2):
        fib[i] = (fib[i - 1] + fib[i - 2]) & modLimit

    for i in range(len(A)):
        # To climb to A[i] rungs, you could either
        # come from A[i]-1 with an one-step jump
        # OR come from A[i]-2 with a two-step jump
        # So from A[i] rungs, the number of different ways of climbing
        # to the top of the ladder is the Fibonacci number at position
        # A[i] + 1
        result[i] = fib[A[i]+1] & ((1<<B[i])-1)
        
    return result


class TestSolution(unittest.TestCase):
    
    def test_solution(self):
        # so = Solution()
        A = [4,4,5,5,1]
        B = [3,2,4,3,1]
        print(solution(A, B))
        
    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')

    def test_isupper(self):
        self.assertTrue('FOO'.isupper())
        self.assertFalse('Foo'.isupper())

    def test_split(self):
        s = 'hello world'
        self.assertEqual(s.split(), ['hello', 'world'])
        # check that s.split fails when the separator is not a string
        with self.assertRaises(TypeError):
            s.split(2)

if __name__ == '__main__':
    unittest.main()

