package main

/*
 * x 的平方根
 * 求平方根算法，本题使用牛顿迭代法
 */
func mySqrt(x int) int {
	return newtonMethod(x)
}

func newtonMethod(x int) int  {
	if x== 0{
		return 0
	}
	var last float64 = 0
	var rest float64 = 1
	for ;last != rest;{
		last = rest
		rest = (rest + float64(x) / rest) / 2
	}
	return int(rest)
}

func MySqrt(x int) int {
	return mySqrt(x)
}


