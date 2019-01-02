package main

/*
 * 最长前缀
 * https://leetcode-cn.com/problems/longest-common-prefix
 *
 */
func longestCommonPrefix(strs []string) string {
	return LongestCommonPrefix(strs)
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}
	var s = make([]byte,0,len(strs[0]))
SEARCHLOOP:
	for i:=0;i<len(strs[0]);i++{
		base := strs[0][i]
		for _,v := range strs[1:]{
			if i == len(v){
				break SEARCHLOOP
			}
			if base != v[i]{
				break SEARCHLOOP
			}
		}
		s = append(s,base)
	}
	return string(s)
}