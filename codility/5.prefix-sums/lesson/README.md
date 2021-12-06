# Theory Implementation


```python
# Counting prefix sums — O(n)
def prefix_sums(A):
    n = len(A)
    P = [0] * (n + 1)
    for k in xrange(1, n + 1):
        P[k] = P[k - 1] + A[k - 1]
    return P

# Total of one slice — O(1)
def count_total(P, x, y):
    return P[y + 1] - P[x]

# Mushroom picker — O(n + m)
def mushrooms(A, k, m):
    n = len(A)
    result = 0
    pref = prefix_sums(A)
    
    for p in xrange(min(m, k) + 1):
        left_pos = k - p
        right_pos = min(n - 1, max(k, k + m - 2 * p))
        result = max(result, count_total(pref, left_pos, right_pos))

    for p in xrange(min(m + 1, n - k)):
        right_pos = k + p
        left_pos = max(0, min(k, k - (m - 2 * p)))
        result = max(result, count_total(pref, left_pos, right_pos))
    return result
```

### Problem: Mushroom Picker
You are given a non-empty, zero-indexed array 
A of n (1 <= n <= 100000) integers a[0], a[1],
..., a[n−1] ( 0 <= a[i] <= 1000 ). 

This array represents number of mushrooms growing
on the consecutive spots along a road. You are also 
given integers k and m (0 <= k, m < n).

A mushroom picker is at spot number k on the road 
and should perform m moves. In one move she moves 
to an adjacent spot. She collects all the mushrooms 
growing on spots she visits. The goal is to calculate 
the maximum number of mushrooms that the mushroom
picker can collect in m moves.
