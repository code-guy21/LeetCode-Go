package twosum

import (
    "sort"
    "reflect"
    "testing"
)

func TestTwoSumVariations(t *testing.T) {
    // Test cases for sorted array variation
	sortedTests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},    
		{[]int{2, 3, 4}, 6, []int{1, 3}},         
		{[]int{-4, -3, -2, -1}, -7, []int{1, 2}}, 
	}
	

    // Test twoSumSorted
    for _, tt := range sortedTests {
        nums := make([]int, len(tt.nums))
        copy(nums, tt.nums)
        sort.Ints(nums)
        got := twoSumSorted(nums, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("twoSumSorted(%v, %v) = %v, want %v", nums, tt.target, got, tt.want)
        }
    }

    // Test cases for other variations
    tests := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
        {[]int{-1, -2, -3, -4}, -7, []int{2, 3}},
    }

    // Test minSpace and optimized variations
    for _, tt := range tests {
        // Test twoSumMinSpace
        got := twoSumMinSpace(tt.nums, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("twoSumMinSpace(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
        }

        // Test twoSumOptimized
        got = twoSumOptimized(tt.nums, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("twoSumOptimized(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
        }
    }
}

// Benchmarks 
func BenchmarkTwoSumVariations(b *testing.B) {
    // Small array
    smallNums := []int{2, 7, 11, 15}
    smallTarget := 9

    // Large array
    largeNums := make([]int, 1000)
    for i := range largeNums {
        largeNums[i] = i
    }
    largeTarget := 1990 // sum of 995 + 995

    b.Run("Original-Small", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSum(smallNums, smallTarget)
        }
    })

    b.Run("Sorted-Small", func(b *testing.B) {
        sorted := make([]int, len(smallNums))
        copy(sorted, smallNums)
        sort.Ints(sorted)
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            twoSumSorted(sorted, smallTarget)
        }
    })

    b.Run("MinSpace-Small", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSumMinSpace(smallNums, smallTarget)
        }
    })

    b.Run("Optimized-Small", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSumOptimized(smallNums, smallTarget)
        }
    })

    b.Run("Original-Large", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSum(largeNums, largeTarget)
        }
    })

    b.Run("Sorted-Large", func(b *testing.B) {
        sorted := make([]int, len(largeNums))
        copy(sorted, largeNums)
        sort.Ints(sorted)
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            twoSumSorted(sorted, largeTarget)
        }
    })

    b.Run("MinSpace-Large", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSumMinSpace(largeNums, largeTarget)
        }
    })

    b.Run("Optimized-Large", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            twoSumOptimized(largeNums, largeTarget)
        }
    })
}