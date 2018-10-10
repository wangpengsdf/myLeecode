package main

/*
 * 实现strStr()
 * https://leetcode-cn.com/problems/implement-strstr
 *
 */

func strStr(haystack string, needle string) int {
	return StrStr(haystack,needle)
}

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	ln := len(needle)
	lh := len(haystack)

	if lh < ln {
		return -1
	}

	flag := 0
	for i := 0; i < lh; i++ {
		if haystack[i] == needle[0] && i+ln-1 < lh {
			for j := 0; j < ln; j++ {
				if haystack[i+j] != needle[j] {
					continue
				} else {
					flag += 1
				}
			}
			if flag == ln {
				return i
			}
			flag = 0
		}
	}
	return -1
}