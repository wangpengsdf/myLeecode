package main

import "testing"

func TestReverse(t *testing.T) {
	t.Log(Reverse(0))
	t.Log(Reverse(1234))
	t.Log(Reverse(-1234))
	t.Log(Reverse(-1230))
}