package main

/*
 * 数组链表去重
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	current := head
	current.Next = nil
	for ;;{
		if next == nil{
			return head
		}
		if next.Val == current.Val{
			next = next.Next
		}else{
			current.Next= next
			current = next
			next = next.Next
			current.Next = nil
		}
	}
}

func DeleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}