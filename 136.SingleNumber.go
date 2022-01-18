package leetcode

func SingleNumber(nums []int) int {
	var single int

	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}
