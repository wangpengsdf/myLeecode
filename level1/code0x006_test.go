package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Log(IsPalindrome(12321))
	t.Log(IsPalindrome(123321))
	t.Log(IsPalindrome(10))
	t.Log(IsPalindrome(-12321))
}
