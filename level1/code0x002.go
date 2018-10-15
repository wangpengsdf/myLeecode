package main

/*
 * 两数之和，
 * https://leetcode-cn.com/problems/two-sum/
 *
 */
func TwoSum(nums []int, target int) []int {
	var sumMap = make(map[int][]int)
	for idx,v :=range nums{
		if _, ok := sumMap[v];ok{
			sumMap[v] = append(sumMap[v],idx)
		}else{
			sumMap[v] = []int{idx}
		}
		if _,ok := sumMap[target-v];ok{
			if target -v != v{
				return []int{sumMap[v][0],sumMap[target-v][0]}
			}else{
				if len(sumMap[v]) > 1{
					return sumMap[v][:2]
				}
			}
		}
	}
	panic("not in list")
}

func twoSum(nums []int, target int) []int {
	return TwoSum(nums,target)
}