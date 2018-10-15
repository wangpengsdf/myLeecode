package main

/*
 * 爬楼梯
 * https://leetcode-cn.com/problems/climbing-stairs/
 */
func climbStairs(n int) int {
	if n == 1{
		return 1
	}
	if n == 2{
		return 2
	}
	var result = make([]int,n+1)
	result[1] = 1
	result[2] = 2
 	for i:= 3;i<=n;i++{
 		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}

func ClimbStairs(n int) int {
	return climbStairs(n)
}