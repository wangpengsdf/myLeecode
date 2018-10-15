package main

/*
 * 二叉树的最小深度
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
 */

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil{
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}else if root.Left != nil && root.Right == nil{
		return minDepth(root.Left) + 1
	}else{
		left := minDepth(root.Left)
		right := minDepth(root.Right)
		if left < right{
			return left + 1
		}else{
			return right + 1
		}
	}
}
