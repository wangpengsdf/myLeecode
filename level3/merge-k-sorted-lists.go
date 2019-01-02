package level3

import(
	. "myLeecode/common"
)

/*
 * 合并N个链表列表
 * https://leetcode-cn.com/problems/merge-k-sorted-lists
 */

// 使用查找法实现
func MergeNLists(nodes... *ListNode) *ListNode {
	var head ,tail *ListNode
	for{
		var foundidx = -1
		var min int
		// 找到最小的点
		for idx,node := range nodes{
			if node != nil{
				if foundidx < 0{
					foundidx = idx
					min = node.Val
				}else{
					if node.Val < min{
						foundidx = idx
						min = node.Val
					}
				}
			}
		}
		// 找不到可以连接的点就可以返回了
		if foundidx < 0{
			return head
		}
		// 最小的向后进一位
		var node = nodes[foundidx]
		nodes[foundidx] = node.Next
		if head == nil{
			// 返回链表初始化
			head = node
			tail = head
			tail.Next = nil
		}else{
			// 链表后移一位
			tail.Next = node //连接当前尾节点和下一个节点
			tail = node // 移动尾节点到最后
		}
	}
}

