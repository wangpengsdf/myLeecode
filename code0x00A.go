package main

/*
 * 数组原地去重
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
 *
 */

func removeDuplicates(nums []int) int {
	return RemoveDuplicates(nums)
}

func RemoveDuplicates(nums []int) int {
	var lastIndex = 0
	for i:=1;i<len(nums);i++{
		if nums[i] == nums[lastIndex]{
			continue
		}else{
			lastIndex += 1
			nums[lastIndex] = nums[i]
		}
	}
	return lastIndex + 1
}