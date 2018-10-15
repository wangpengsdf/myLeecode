package main

/*
 * 杨辉三角
 * https://leetcode-cn.com/problems/pascals-triangle/
 */
func generate(numRows int) [][]int {
	if numRows == 0{
		return [][]int{}
	}
	var lines = [][]int{{1}}
	for n := 1;n<numRows;n++{
		var line = make([]int,0,len(lines[n-1])+1)
		line = append(line,1)
		for i:=0;i<len(lines[n-1]) -1;i++{
			line = append(line,lines[n-1][i] + lines[n-1][i+1])
		}
		line = append(line,1)
		lines = append(lines,line)
	}
	return lines
}

