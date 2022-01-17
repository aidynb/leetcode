package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if value, exists := m[target-num]; exists {
			return []int{value, i}
		} else {
			m[num] = i
		}
	}
	return nil
}
