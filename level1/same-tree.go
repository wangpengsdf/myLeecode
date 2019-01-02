package main

import . "myLeecode/common"
/*
 * 相同的树
 * https://leetcode-cn.com/problems/same-tree/
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{return true}
	if p != nil && q == nil{return false}
	if p == nil && q != nil{return false}
	if p.Val != q.Val{return false}
	return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTree(p,q)
}

