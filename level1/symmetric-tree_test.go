package main

import . "myLeecode/common"

import "testing"

func TestIsSymmetric(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 3
	t.Log(isSymmetric(root))
}