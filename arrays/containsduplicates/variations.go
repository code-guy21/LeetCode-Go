package containsduplicate

import "sort"

// Sorting approach
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

// BitSet approach for small positive integers
func containsDuplicateBitSet(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }

    // Find the maximum number in nums to allocate a sufficiently large bitSet
    maxNum := 0
    for _, num := range nums {
        if num < 0 {
            // Fall back to map-based approach for negative numbers
            return containsDuplicate(nums)
        }
        if num > maxNum {
            maxNum = num
        }
    }

    // Dynamically allocate bitSet for range 0 to maxNum
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


// Two pointers approach for sorted input
func containsDuplicateTwoPointers(nums []int) bool {
   if len(nums) <= 1 {
       return false
   }
   
   sorted := make([]int, len(nums))
   copy(sorted, nums)
   sort.Ints(sorted)
   
   left := 0
   for right := 1; right < len(sorted); right++ {
       if sorted[left] == sorted[right] {
           return true
       }
       left = right
   }
   return false
}