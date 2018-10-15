package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(LongestCommonPrefix([]string{"flower","flow","flight"}))
	t.Log(LongestCommonPrefix([]string{"dog","racecar","car"}))
}