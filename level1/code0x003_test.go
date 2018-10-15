package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1 = new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3
	var l2 = new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	t.Log(addTwoNumbers(l1,l2))
}
