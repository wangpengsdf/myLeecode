package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var s = []int{1,1,2,3,3}
	t.Log(RemoveDuplicates(s))
	t.Log(s)
}