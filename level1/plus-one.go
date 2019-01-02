package main

/*
 * 数组加1
 * https://leetcode-cn.com/problems/plus-one/
 */
func plusOne(digits []int) []int {
	var extra = 1
	for i := len(digits)-1;i>=0;i--{
		digits[i] += extra
		extra = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if extra > 0 {
		var s = make([]int,0,len(digits)+1)
		s = append(s,extra)
		s = append(s,digits...)
		return s
	}else{
		return digits
	}
}

func PlusOne(digits []int) []int {
	return plusOne(digits)
}



