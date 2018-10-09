package main

/*
 * 两数相加
 * https://leetcode-cn.com/problems/add-two-numbers
 *
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return AddTwoNumbers(l1,l2)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var extra = 0
	var current *ListNode
	var head *ListNode
	for{
		var sum = 0
		var isNil = true
		if l1 != nil{
			isNil = false
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			isNil = false
			sum += l2.Val
			l2 = l2.Next
		}
		if extra != 0{
			sum += extra
			isNil = false
		}
		if isNil{
			return head
		}

		extra = sum / 10
		var left = sum % 10
		if current == nil{
			current = new(ListNode)
			current.Val = left
			head = current
		}else{
			current.Next = new(ListNode)
			current = current.Next
			current.Val = left
		}

	}
}