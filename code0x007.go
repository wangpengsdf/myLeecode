package main

/*
 * 罗马数字转整数
 * https://leetcode-cn.com/problems/roman-to-integer
 *
 */

 var baseNum = map[rune]int{
 	'I':1,
 	'V':5,
 	'X':10,
 	'L':50,
 	'C':100,
 	'D':500,
 	'M':1000,
 }

func romanToInt(s string) int {
	return RomanToInt(s)
}

func RomanToInt(s string) int {
	var total = 0
	var runes = []rune(s)
	for i:=0;i<len(runes);{
		if i+1 < len(runes){
			isSpecial,num := isSpecial(runes[i],runes[i+1])
			if isSpecial{
				total += num
				i += 2
				continue
			}
		}
		total += baseNum[runes[i]]
		i += 1
	}
	return total
}

func isSpecial(front,end rune)(bool,int) {
	var isSpecial = false
	if front == 'I'{
		if end == 'V' || end == 'X'{
			isSpecial = true
		}
	}else if front == 'X'{
		if end == 'L' || end == 'C'{
			isSpecial = true
		}
	}else if front == 'C'{
		if end == 'D' || end == 'M'{
			isSpecial = true
		}
	}
	if isSpecial{
		return true, baseNum[end] -baseNum[front]
	}else{
		return false, 0
	}
}