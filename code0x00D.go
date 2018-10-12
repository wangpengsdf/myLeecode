package main

import (
	"strconv"
)

/*
 * 报数
 * https://leetcode-cn.com/problems/count-and-say
 */

func countAndSay(n int) string{
	if n == 1{
		return "1"
	}
	var prestr = []rune(CountAndSay(n-1))
	var newstr = ""
	var currentRune rune = -1
	var count = 0
	for _,scanRune := range prestr {
		if scanRune != currentRune{
			if count > 0 {
				newstr += strconv.Itoa(count) + string(currentRune)
			}
			currentRune = scanRune
			count = 1
		}else{
			count += 1
		}
	}
	newstr += strconv.Itoa(count) + string(currentRune)
	return string(newstr)
}

func CountAndSay(n int) string {
	return countAndSay(n)
}