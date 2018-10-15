package main

/*
 * 二进制求和
 * https://leetcode-cn.com/problems/add-binary
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	reverseTree(root.Left)
	return isSameTree(root.Left,root.Right)
}

func reverseTree(root *TreeNode)  {
	if root == nil{
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	reverseTree(root.Left)
	reverseTree(root.Right)
}

func IsSymmetric(root *TreeNode) bool {
	return isSymmetric(root)
}