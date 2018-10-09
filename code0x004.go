package main

import "strings"

/*
 * 通配符
 * https://leetcode-cn.com/problems/wildcard-matching/
 *
 */

func isMatch(s string, p string) bool {
	return IsMatch(s,p)
}

func IsMatch(s string, p string) bool {
	var fragments = strings.Split(p,"*")
	
}

func charEqual(s,p rune) bool {
	return s == p || p == '?'
}