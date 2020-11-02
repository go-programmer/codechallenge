
# https://docs.python.org/3/library/unittest.html

import unittest



def solution(A):
    A.sort()        # Sort A in non-decreasing order
    if A[0] >= 0:   return A[0] + A[0]      # All non-negative
    if A[-1] <= 0:  return -A[-1] - A[-1]   # All non-positive
    front, back = len(A)-1, 0
    minAbs = A[-1] + A[-1]                  # Final result

    # Travel the array from both ends to some center point.
    # See following post for the proof of this method.
    # https://codesays.com/2014/solution-to-min-abs-sum-of-two-by-codility
    while back <= front:
        temp = abs(A[back] + A[front])

        # Update the result if needed
        if temp < minAbs:  minAbs = temp
        
        # Adjust the pointer for next 
        if back == front: 
            break
        elif abs(A[back+1] + A[front]) <= temp:     
            back += 1
        elif abs(A[back] + A[front-1]) <= temp:     
            front -= 1
        else:                                       
            back += 1
            front -= 1

    return minAbs



class TestSolution(unittest.TestCase):

    def test_solution(self):
        A = [1,4,-3]
        want = 1    
        self.assertEqual(solution(A), want)

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
