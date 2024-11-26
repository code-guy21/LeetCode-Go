# Two Sum - Problem Solution

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers in `nums` such that they add up to `target`.

### Constraints:

- Each input has exactly one solution
- You may not use the same element twice
- Return the answer in any order

## Solution Implementations

### 1. Hash Map Solution (Original)

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

### 2. Sorted Array Solution

```go
func twoSumSorted(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1}
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

### 3. Memory Optimized Solution

```go
func twoSumMinSpace(nums []int, target int) []int {
    // Initialize an empty map to store number and its index
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, exists := numMap[complement]; exists {
            result := []int{j, i}
            sort.Ints(result)
            return result
        }
        numMap[num] = i
    }
    return nil
}
```

## Time & Space Complexity

### Hash Map Solution (Original)

- Time Complexity: O(n)
- Space Complexity: O(n)

### Sorted Array Solution

- Time Complexity: O(n log n) due to sorting
- Space Complexity: O(1)

### Memory Optimized Solution

- Time Complexity: O(n log n)
- Space Complexity: O(n)

## Running Tests

### Basic Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

### Run Specific Tests

```bash
# Run only TwoSum tests
go test -run TestTwoSum

# Run only variation tests
go test -run TestTwoSumVariations
```

### Performance Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run benchmarks with memory allocation statistics
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkTwoSum
```

## Test Files Structure

```
twosum/
├── solution.go        # Original solution
├── solution_test.go   # Original tests
├── variations.go      # Additional implementations
└── variations_test.go # Tests for variations
```

## Edge Cases Covered

- Empty array
- Array with only one element
- Negative numbers
- Duplicate numbers
- Target sum with same number

## Interview Tips

1. Always clarify constraints first
2. Discuss brute force approach before optimization
3. Explain space-time trade-offs
4. Consider follow-up questions about scaling or constraints

## Common Patterns Used

- Hash Table: For quick lookups to find complements.
- Two Pointers (Sorted Variant): Efficiently finding pairs in a sorted array.
- Array Traversal: Iterating through the array to process elements.

## Related Problems

- Three Sum: Find three numbers in an array that add up to a target.
- Two Sum II (sorted array): Similar to Two Sum but the input array is already sorted.
- Two Sum III (data structure design): Design a data structure that supports adding numbers and finding pairs that sum to a target.
