package lc

func searchInsertPosition(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		}
	}

	if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}
