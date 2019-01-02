package main

import . "myLeecode/common"


/*
 * 二叉树深度判断
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 */
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right{
		return left + 1
	}else{
		return right + 1
	}
}
