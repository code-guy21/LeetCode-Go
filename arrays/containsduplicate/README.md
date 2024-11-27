# Contains Duplicate

## Problem Description

Given an integer array `nums`, return `true` if any value appears at least twice in the array, and `false` if every element is distinct.

### Examples

```
Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false
```

## Solutions

### 1. Hash Map Approach (Original)

```go
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

    for _, num := range nums {
        if _, exists := seen[num]; exists {
            return true
        }
        seen[num] = struct{}{}
    }

    return false
}
```

**Complexity:**

- Time: O(n)
- Space: O(n)

### 2. Sorting Approach

```go
func containsDuplicateSort(nums []int) bool {
   if len(nums) <= 1 {
       return false
   }

   sorted := make([]int, len(nums))
   copy(sorted, nums)
   sort.Ints(sorted)

   for i := 1; i < len(sorted); i++ {
       if sorted[i] == sorted[i-1] {
           return true
       }
   }
   return false
}
```

**Complexity:**

- Time: O(n log n)
- Space: O(1) if in-place sorting, otherwise O(n)

### 3. BitSet Approach (for positive integers)

```go
func containsDuplicateBitSet(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }

    maxNum := 0
    for _, num := range nums {
        if num < 0 {
            return containsDuplicate(nums) // Fallback for negatives
        }
        if num > maxNum {
            maxNum = num
        }
    }

    bitSet := make([]uint64, maxNum/64+1)

    for _, num := range nums {
        bucket := num / 64
        bit := uint64(1) << uint(num%64)

        if bitSet[bucket]&bit != 0 {
            return true
        }
        bitSet[bucket] |= bit
    }
    return false
}
```

**Complexity:**

- Time: O(n)
- Space: O(1) for limited range of integers

## Running Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run benchmarks
go test -bench=. -benchmem
```

## Implementation Notes

### Hash Map Approach

- Best for general cases
- Handles any integer values
- Constant time lookups

### Sorting Approach

- No extra space needed (if sorting is in-place)
- Good for memory-constrained systems
- Original array order is modified

### BitSet Approach

- Extremely memory efficient for small positive integers
- Falls back to Hash Map for negatives
- Limited by integer range

## Edge Cases

- Empty array
- Single element
- All duplicates
- No duplicates
- Negative numbers
- Large numbers

## Benchmarking Results

| Method   | Time Complexity | Space Complexity | Notes                                     |
| -------- | --------------- | ---------------- | ----------------------------------------- |
| Hash Map | O(n)            | O(n)             | General-purpose, works with all integers  |
| Sorting  | O(n log n)      | O(1) / O(n)      | Memory-efficient if in-place sorting used |
| BitSet   | O(n)            | O(1)             | Optimized for small positive integers     |

## Related Problems

- Contains Duplicate II
- Contains Duplicate III
- Find All Duplicates in an Array
