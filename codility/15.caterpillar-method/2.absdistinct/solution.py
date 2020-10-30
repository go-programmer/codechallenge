
# https://docs.python.org/3/library/unittest.html

import unittest
def solution(A):
    return len(set([abs(x) for x in A]))

class TestSolution(unittest.TestCase):

    def test_solution(self):
        pythonSet = {"apple", "banana", "cherry","banana"}
        print(pythonSet) 
        thisset = set(("apple", "banana", "cherry")) # note the double round-brackets
        print(thisset) 

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
