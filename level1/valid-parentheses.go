package main

/*
 * 括号有效性判断
 * https://leetcode-cn.com/problems/valid-parentheses
 *
 */

var parentTheses = map[rune]rune{
	'(':')',
	'[':']',
	'{':'}',
}

func isValid(s string) bool {
	return IsValid(s)
}

func IsValid(s string) bool {
	var lastRune []rune
	for _,v := range []rune(s){
		if _,ok := parentTheses[v];ok{
			lastRune = append(lastRune,v)
		}else{
			if len(lastRune) == 0{
				return false
			}
			head := lastRune[len(lastRune)-1]
			if parentTheses[head] != v{
				return false
			}else{
				lastRune = lastRune[0:len(lastRune)-1]
			}
		}
	}
	return len(lastRune) == 0
}
