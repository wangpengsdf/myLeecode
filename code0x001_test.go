package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var l1 = &ListNode{Val:1}
	l1.Next =&ListNode{Val:9}
	var l2 = &ListNode{Val:3}
	l2.Next =&ListNode{Val:6}
	t.Log(MergeNLists(l1,l2))
	var l3 *ListNode
	var l4 = &ListNode{Val:0}
	t.Log(MergeTwoLists(l3,l4))
}

func TestMergeNLists(t *testing.T) {
	var l1 = &ListNode{Val:1}
	l1.Next =&ListNode{Val:9}
	var l2 = &ListNode{Val:3}
	l2.Next =&ListNode{Val:6}
	var l3 = &ListNode{Val:2}
	l3.Next =&ListNode{Val:7}
	t.Log(MergeNLists(l1,l2,l3))
}
