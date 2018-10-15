package main

/*
 * 回文数判断
 * https://leetcode-cn.com/problems/palindrome-number
 */

func isPalindrome(x int) bool {
	return IsPalindrome(x)
}

func IsPalindrome(x int) bool {
	if x < 0{
		return false
	}
	if x % 10 == 0 && x != 0{
		return false
	}
	var nums = make([]int,0,16)
	for{
		if x == 0{
			break
		}
		nums = append(nums, x % 10)
		x = x / 10
	}
	length := len(nums)
	for i:=0;i<length/2;i++{
		if nums[i] != nums[length - 1 - i]{
			return false
		}
	}
	return true
}