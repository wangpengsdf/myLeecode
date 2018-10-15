package main

import "testing"

func TestMySqrt(t *testing.T) {
	t.Log(MySqrt(5))
	t.Log(MySqrt(64))
	t.Log(MySqrt(12039120974981))
}