package main

/*
 * 移除数组指定元素
 * https://leetcode-cn.com/problems/remove-element/
 */


func removeElement(nums []int, val int) int {
	return RemoveElement(nums,val)
}

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0{
		return 0
	}
	readindex := 0
	for i:=0;i<len(nums);i++{
		if nums[i] != val{
			nums[readindex] = nums[i]
			readindex += 1
		}else{
			continue
		}
	}
	return readindex
}