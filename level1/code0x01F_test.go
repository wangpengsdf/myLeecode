package main

import "testing"

func TestGeneratePascals(t *testing.T) {
	for _,v := range GeneratePascals(100){
		t.Log(v)
	}
}