package main

/*
 * 二进制求和
 * https://leetcode-cn.com/problems/add-binary
 */
func addBinary(a string, b string) string {
	if len(a) > len(b){
		a, b = b, a
	}
	astr := []rune(a)
	bstr := []rune(b)
	var extra = 0
	for i:= 0;i<len(bstr);i++{
		aidx := len(astr) - i - 1
		bidx := len(bstr) - i - 1
		sum := 0
		if aidx >= 0{
			if astr[aidx] == '1'{
				sum += 1
			}
		}
		if bidx >= 0 {
			if bstr[bidx] == '1'{
				sum += 1
			}
		}
		sum += extra
		left := sum % 2
		extra = sum / 2
		if left == 1{
			bstr[bidx] = '1'
		}else{
			bstr[bidx] = '0'
		}
	}
	if extra > 0{
		var str = make([]rune,0,len(bstr) + 1)
		str = append(str,'1')
		return string(append(str, bstr...))
	}else{
		return string(bstr)
	}
}

func AddBinary(a string, b string) string {
	return addBinary(a,b)
}


