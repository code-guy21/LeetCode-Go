# Two Sum - Problem Solution

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers in `nums` such that they add up to `target`.

### Constraints:

- Each input has exactly one solution
- You may not use the same element twice
- Return the answer in any order

## Approach 1: Hash Map (Optimal)

### Time Complexity: O(n)

### Space Complexity: O(n)

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if j, exists := seen[target-num]; exists {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}
```

### Step-by-Step Explanation:

1. Create a hash map to store numbers we've seen and their indices
2. For each number:
   - Check if its complement (target - current number) exists in the map
   - If found, return current index and stored index
   - If not found, store current number and index
3. Return nil if no solution found

### Performance Analysis:

- Single pass through the array
- Constant-time lookups in hash map
- Space grows linearly with input size

## Approach 2: Two Pointers (Requires Sorting)

### Time Complexity: O(n log n)

### Space Complexity: O(1)

```go
func twoSumSorted(nums []int, target int) []int {
    // Create index pairs to track original indices
    pairs := make([][2]int, len(nums))
    for i, num := range nums {
        pairs[i] = [2]int{num, i}
    }

    // Sort by values
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][0] < pairs[j][0]
    })

    // Two pointers approach
    left, right := 0, len(nums)-1
    for left < right {
        sum := pairs[left][0] + pairs[right][0]
        if sum == target {
            return []int{pairs[left][1], pairs[right][1]}
        }
        if sum < target {
            left++
        } else {
            right--
        }
    }
    return nil
}
```

### Benchmarks

```go
func BenchmarkTwoSum(b *testing.B) {
    nums := []int{2, 7, 11, 15}
    target := 9
    for i := 0; i < b.N; i++ {
        twoSum(nums, target)
    }
}
```

## Edge Cases:

- Empty array
- Array with only one element
- Negative numbers
- Multiple valid pairs (problem guarantees unique solution)
- No valid solution

## Test Cases:

```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
        {[]int{1}, 1, nil},
        {[]int{}, 0, nil},
    }

    for _, tt := range tests {
        got := twoSum(tt.nums, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("twoSum(%v, %v) = %v, want %v",
                     tt.nums, tt.target, got, tt.want)
        }
    }
}
```

## Interview Tips:

1. Always clarify constraints first
2. Discuss brute force approach before optimization
3. Explain space-time trade-offs
4. Consider follow-up questions about scaling or constraints

## Common Patterns Used:

- Hash Table
- Two Pointers (sorted variant)
- Array Traversal

## Related Problems:

- Three Sum
- Two Sum II (sorted array)
- Two Sum III (data structure design)
