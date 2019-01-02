package main

import . "myLeecode/common"

/*
 * 二叉树的层次遍历
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
 */

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	var queue = []*TreeNode{root}
	var element [][]int
	for {
		var newqueque = make([]*TreeNode,0,len(queue)*2)
		var line = make([]int,0,len(queue)*2)
		for _,v := range queue{
			line = append(line, v.Val)
			if v.Left != nil{
				newqueque = append(newqueque,v.Left)
			}
			if v.Right != nil{
				newqueque = append(newqueque,v.Right)
			}
		}
		element = append(element,line)
		if len(newqueque) == 0{
			break
		}
		queue = newqueque
	}
	for i:=0;i<len(element) /2 ;i++{
		element[i],element[len(element) -i -1] = element[len(element) -i -1], element[i]
	}
	return element
}