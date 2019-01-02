package main

import . "myLeecode/common"

/*
 * 路径总和
 * https://leetcode-cn.com/problems/path-sum/
 */


func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	extra := sum - root.Val
	if extra == 0 && root.Left == nil && root.Right == nil{
		return true
	}
	return hasPathSum(root.Left,extra) ||  hasPathSum(root.Right,extra)
}