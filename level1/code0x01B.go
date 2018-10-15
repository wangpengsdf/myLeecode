package main

/*
 * 二进制求和
 * https://leetcode-cn.com/problems/add-binary
 */

func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	var depthDelta = maxDepth(root.Left) - maxDepth(root.Right)
	if depthDelta > 1 || depthDelta < -1{
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}


