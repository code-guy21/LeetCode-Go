package twosum

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