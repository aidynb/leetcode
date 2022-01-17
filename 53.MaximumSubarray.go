package leetcode

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	var tempSum int

	for _, currNum := range nums {
		tempSum += currNum

		if tempSum > maxSum {
			maxSum = tempSum
		}

		if tempSum < 0 {
			tempSum = 0
		}
	}
	return maxSum
}
