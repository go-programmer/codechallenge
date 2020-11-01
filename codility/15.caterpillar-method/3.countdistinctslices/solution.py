
# https://docs.python.org/3/library/unittest.html

import unittest

def martinkyselSolution(M, A):
    the_sum = 0
    front = back = 0
    seen = [False] * (M+1)
    while (front < len(A) and back < len(A)):
        while (front < len(A) and seen[A[front]] != True):
            the_sum += (front-back+1)
            seen[A[front]] = True
            front += 1
        else:
            while front < len(A) and back < len(A) and A[back] != A[front]:
                seen[A[back]] = False
                back += 1
                 
            seen[A[back]] = False
            back += 1
    return min(the_sum, 1000000000)

def codesaysSolution(M, A):
    accessed = [-1] * (M + 1)   # -1: not accessed before
                                # Non-negative: the previous occurrence position
    front, back = 0, 0
    result = 0
    for front in range(len(A)):
        if accessed[A[front]] == -1:
            # Met with a new unique item
            accessed[A[front]] = front
        else:
            # Met with a duplicate item.
            # Compute the number of distinct slices 
            # between newBack-1 and back position.
            newBack = accessed[A[front]] + 1
            
            # result += (newBack - back) * (front - back + front - newBack + 1) / 2
            
            result += (front - back) * (front - back + 1) / 2
            result -= (front - newBack) * (front - newBack + 1) / 2

            if result >= 1000000000:   return 1000000000

            # Restore and set the accessed array
            for index in range(back, newBack):
                accessed[A[index]] = -1
                
            accessed[A[front]] = front
            back = newBack

    # Process the last slices
    result += (front - back + 1) * (front - back + 2) / 2
    return min(int(result), 1000000000)

class TestSolution(unittest.TestCase):

    def xtest_martinkyselSolution(self):
        A = [3,4,5,5,2]
        M = 6
        want = 9    
        self.assertEqual(martinkyselSolution(M, A), want)

    def test_codesaysSolution(self):
        A = [3,4,5,5,2]
        M = 6
        want = 9    
        self.assertEqual(codesaysSolution(M, A), want)

    # def test_isupper(self):
    #     self.assertTrue('FOO'.isupper())
    #     self.assertFalse('Foo'.isupper())

    # def test_split(self):
    #     s = 'hello world'
    #     self.assertEqual(s.split(), ['hello', 'world'])
    #     # check that s.split fails when the separator is not a string
    #     with self.assertRaises(TypeError):
    #         s.split(2)

if __name__ == '__main__':
    unittest.main()
