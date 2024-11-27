package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T){
	tests := []struct{
		nums []int
		want bool
	}{
	{[]int{1, 2, 3, 1}, true},
        {[]int{1, 2, 3, 4}, false},
        {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
        {[]int{}, false},
        {[]int{1}, false},
        {[]int{1, 1}, true},
	}

	for _, tt := range tests {
		got := containsDuplicate(tt.nums)

		if got != tt.want {
			t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func BenchmarkContainsDuplicate(b *testing.B){
	nums := []int{1,2,3,4,5,6,7,8,9,1}

	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}
