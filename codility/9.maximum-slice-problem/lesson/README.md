# Theory Implementation
Maximum slice can be solved with various time complexities, from naive to O(N).

#### O(n 3 ) solution
```python
def slow_max_slice(A):
    n = len(A)
    result = 0
    for p in xrange(n):
        for q in xrange(p, n):
            sum = 0
            for i in xrange(p, q + 1):
                sum += A[i]
            result = max(result, sum)
    return result
```


### O(n^2) solutiong using prefix-sum
```python
def quadratic_max_slice(A, pref):
    n = len(A), result = 0
    for p in xrange(n):
        for q in xrange(p, n):
            sum = pref[q + 1] - pref[p]
            result = max(result, sum)
    return result
```

### O(n^2) solution without prefix-sum
```python
def quadratic_max_slice(A):
    n = len(A), result = 0
    for p in xrange(n):
        sum = 0
        for q in xrange(p, n):
            sum += A[q]
            result = max(result, sum)
    return result
```

### O(n) solution
```python
def golden_max_slice(A):
    max_ending = max_slice = 0
    for a in A:
        max_ending = max(0, max_ending + a)
        max_slice = max(max_slice, max_ending)
    return max_slice
```
For each position, we compute the largest sum that
ends in that position. If we assume that the maximum
sum of a slice ending in position i equals max_ending,
then the maximum slice ending in position i+1 equals
max(0, max_ending + a[i+1]).