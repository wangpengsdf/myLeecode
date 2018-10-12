package main

import "testing"

func TestCountAndSay(t *testing.T) {
	for i := 1; i <= 30;i++{
		t.Log(CountAndSay(i))
	}
}