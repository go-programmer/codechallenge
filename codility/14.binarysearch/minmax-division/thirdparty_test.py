
import unittest

def blockSizeIsValid(A, max_block_cnt, max_block_size):
    block_sum = 0
    block_cnt = 0
 
    for element in A:
        if block_sum + element > max_block_size:
            block_sum = element
            block_cnt += 1
        else:
            block_sum += element
        
        if block_cnt >= max_block_cnt:
            return False
 
    return True
 
def binarySearch(A, max_block_cnt, using_M_will_give_you_wrong_results):
    lower_bound = max(A)
    upper_bound = sum(A)
 
    if max_block_cnt == 1:      return upper_bound
    if max_block_cnt >= len(A): return lower_bound
 
    while lower_bound <= upper_bound:
        candidate_mid = (lower_bound + upper_bound) // 2
        ret = blockSizeIsValid(A, max_block_cnt, candidate_mid)    
        
        if ret:
            upper_bound = candidate_mid - 1
        else:
            lower_bound = candidate_mid + 1
 
    return lower_bound
 
def solution():
    A=[2, 1, 5, 1, 2, 2, 2]
    K=3
    M=5
    return binarySearch(A,K,M)

class TestSolution(unittest.TestCase):
    
    def test_solution(self):
        
        self.assertEqual(6, solution())

if __name__ == '__main__':
    unittest.main()
