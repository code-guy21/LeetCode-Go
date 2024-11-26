package twosum

import "sort"

// Two Sum II - Input Array Is Sorted
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

// Two Sum with Minimal Space
func twoSumMinSpace(nums []int, target int) []int {
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


// Performance Optimized Two Sum
func twoSumOptimized(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        if j, exists := numMap[target-num]; exists {
            // Ensure indices are in ascending order
            if j < i {
                return []int{j, i}
            }
            return []int{i, j}
        }
        numMap[num] = i
    }
    return nil
}
