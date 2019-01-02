package common


import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//格式化输出链表的内容
func (l *ListNode)String() string{
	var s = fmt.Sprintf("List: %d",l.Val)
	tail := l.Next
	for {
		if tail != nil{
			s += fmt.Sprintf(" %d",tail.Val)
			tail = tail.Next
		}else{
			return s
		}
	}
}
