
# https://docs.python.org/3/library/unittest.html

import unittest

# Check whether a sequence contains a contiguous
# subse-squence whose sum of elements equals s. 
# For example, in the following sequence we are looking
# for a subsequence whose total equals s = 12.
# O(n) time complexity
def caterpillar(A, s):
    n = len(A)
    front, total = 0, 0
    for back in range(n):
        while (front < n and total + A[front] <= s):
            total += A[front]
            front += 1
        if total == s:
            return True
        total -= A[back]
    return False


# Problem: You are given n sticks (of lengths >= 1 and <= 10^9).
# The goal is to count the number of triangles
# that can be constructed using these sticks.

# More precisely, count the number of triplets
# at indices x < y < z, such that A[x] + A[y] > A[z].

# Solution O(n^2):
# For every pair x, y we can find the largest
# stick z that can be used to construct the
# triangle. Every stick k, such that y < k <= z,
# can also be used, because the condition 
# A[x] + A[y] > A[k] will still be true. 
# We can add up all these triangles at once.

# If the value z is found every time from the
# beginning then we get a O(n^3) time complexity
# solution. However, we can instead use the 
# caterpillar method. When increasing the
# value of y, we can increase (as far as possible)
# the value of z.

def triangles(A):
    n = len(A)
    result = 0
    for x in range(n):
        z = x + 2
        for y in range(x + 1, n):
            while (z < n and A[x] + A[y] > A[z]):
                z += 1
            result += z - y - 1
    return result


class TestSolution(unittest.TestCase):

    def test_triangles(self):
        A = [2, 3, 4]
        self.assertEqual(triangles(A),1)
        
        A = [6, 2, 7, 4, 1, 3, 6]
        self.assertEqual(triangles(A),33)
        # A = [1,2,3,4,6,6,7]
        # Function returns differs when sorted.

    def xtest_caterpillar(self):
        A = [6, 2, 7, 4, 1, 3, 6]
        self.assertTrue(caterpillar(A, 12))

    # def test_solution(self):
    #     # so = Solution()
    #     A = [4,4,5,5,1]
    #     B = [3,2,4,3,1]
    #     print(solution(A, B))
        
    # def test_upper(self):
    #     self.assertEqual('foo'.upper(), 'FOO')

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
