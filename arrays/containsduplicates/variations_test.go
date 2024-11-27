package containsduplicate

import "testing"

func TestContainsDuplicateVariations(t *testing.T) {
   tests := []struct {
       name string
       nums []int
       want bool
   }{
       {"basic duplicate", []int{1, 2, 3, 1}, true},
       {"no duplicate", []int{1, 2, 3, 4}, false},
       {"multiple duplicates", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
       {"empty array", []int{}, false},
       {"single element", []int{1}, false},
       {"adjacent duplicates", []int{1, 1}, true},
       {"large numbers", []int{100000, 200000, 100000}, true},
       {"negative numbers", []int{-1, -1, 2}, true},
   }

   for _, tt := range tests {
       t.Run(tt.name, func(t *testing.T) {
           // Test sorting approach
           got := containsDuplicateSort(tt.nums)
           if got != tt.want {
               t.Errorf("containsDuplicateSort(%v) = %v, want %v", tt.nums, got, tt.want)
           }

           // Test two pointers approach
           got = containsDuplicateTwoPointers(tt.nums)
           if got != tt.want {
               t.Errorf("containsDuplicateTwoPointers(%v) = %v, want %v", tt.nums, got, tt.want)
           }

           // Test BitSet approach only for positive numbers
           if allPositive(tt.nums) {
               got = containsDuplicateBitSet(tt.nums)
               if got != tt.want {
                   t.Errorf("containsDuplicateBitSet(%v) = %v, want %v", tt.nums, got, tt.want)
               }
           }
       })
   }
}

func allPositive(nums []int) bool {
   for _, n := range nums {
       if n < 0 {
           return false
       }
   }
   return true
}

func BenchmarkContainsDuplicateVariations(b *testing.B) {
   testCases := []struct {
       name string
       nums []int
   }{
       {"small array", []int{1, 2, 3, 4, 5, 1}},
       {"medium array", generateArray(100, true)},
       {"large array", generateArray(1000, false)},
       {"all positive", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
   }

   for _, tc := range testCases {
       b.Run("Map-"+tc.name, func(b *testing.B) {
           for i := 0; i < b.N; i++ {
               containsDuplicate(tc.nums)
           }
       })

       b.Run("Sort-"+tc.name, func(b *testing.B) {
           for i := 0; i < b.N; i++ {
               containsDuplicateSort(tc.nums)
           }
       })

       b.Run("TwoPointers-"+tc.name, func(b *testing.B) {
           for i := 0; i < b.N; i++ {
               containsDuplicateTwoPointers(tc.nums)
           }
       })

       if allPositive(tc.nums) {
           b.Run("BitSet-"+tc.name, func(b *testing.B) {
               for i := 0; i < b.N; i++ {
                   containsDuplicateBitSet(tc.nums)
               }
           })
       }
   }
}

func generateArray(size int, hasDuplicate bool) []int {
   arr := make([]int, size)
   for i := 0; i < size; i++ {
       arr[i] = i
   }
   if hasDuplicate {
       arr[size-1] = arr[0]
   }
   return arr
}