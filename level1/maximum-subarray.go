package main

func maxSubArray(nums []int) int {
	var sum = nums[0]
	var n = nums[0]
	for i := 1; i < len(nums); i++ {
		if n > 0 {
			n += nums[i]
		} else {
			n = nums[i]
		}
		if sum < n {
			sum = n
		}
	}
	return sum
}

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}