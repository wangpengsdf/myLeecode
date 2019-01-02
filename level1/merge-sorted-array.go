package main

import "sort"

/*
 * 合并有序数组
 * https://leetcode-cn.com/problems/merge-sorted-array/
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	sort.Ints(append(nums1[:m],nums2[:n]...))
}

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	merge(nums1,m,nums2,n)
}

