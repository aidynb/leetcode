package leetcode

func MajorityElement(nums []int) int {
	res := nums[0]
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count = 0
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}
