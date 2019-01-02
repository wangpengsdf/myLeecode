package main

import (
	. "myLeecode/common"
	"myLeecode/level3"
)

/*
 * 主要是这两道题，合并两个列表
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */


//利用合并N个List实现
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return level3.MergeNLists(l1,l2)
}



