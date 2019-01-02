package main

import . "myLeecode/common"

/*
 * 将有序数组转换为二叉搜索树
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
 */


 func sortedArrayToBST(nums []int) *TreeNode {
 	if len(nums) == 0{
 		return nil
	}
	var center = len(nums)/2
	root := new(TreeNode)
	root.Val = nums[center]
	root.Left = sortedArrayToBST(nums[:center])
	root.Right = sortedArrayToBST(nums[center+1:])
	return root
 }


