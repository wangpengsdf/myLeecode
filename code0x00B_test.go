package main

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{1,2,2,3,4}
	t.Log(RemoveElement(nums,2))
	t.Log(nums)
}