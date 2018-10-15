package main

import "math"

/*
 * 反转整数
 * https://leetcode-cn.com/problems/reverse-integer/
 */

func reverse(x int) int {
	return Reverse(x)
}

func Reverse(x int) int {
	var i = 0
	for {
		if x == 0{
			break
		}
		var left = x % 10
		i = i * 10 + left
		if i > math.MaxInt32 || i < math.MinInt32{
			return 0
		}
		x = x / 10
	}
	return i
}