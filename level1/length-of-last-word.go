package main

/*
 * 最后一个单词的长度
 * https://leetcode-cn.com/problems/maximum-subarray/
 */

func lengthOfLastWord(s string) int {
	var prestr = []rune(s)
	var wordlen = 0
	var foundword = false
	for i:= len(prestr);i>0;i--{
		if prestr[i-1] == ' '{
			if foundword{
				break
			}else{
				continue
			}

		}else{
			foundword = true
			wordlen += 1
		}
	}
	return wordlen
}

func LengthOfLastWord(s string) int {
	return lengthOfLastWord(s)
}