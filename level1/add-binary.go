package main

import . "myLeecode/common"

/*
 * 平衡二叉树
 * https://leetcode-cn.com/problems/balanced-binary-tree/
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


