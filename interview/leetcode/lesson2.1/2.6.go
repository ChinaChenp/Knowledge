package main

import "fmt"

//Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
//For example, Given [100, 4, 200, 1, 3, 2], The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.
//给定未排序的整数数组，找到最长连续元素序列的长度。
//例如，给定[100,4,2,1,3,2]，最长的连续元素序列是[1,2,3,4]。 返回它的长度：4。你的算法应该以O（n）复杂度运行。


// 思路
// 用一个map[int]bool 记录每个元素的使用情况，对每个元素以该元素为中心，往左右两边扩张，直到不连续为止
func longestConsecutive(arr []int) int {
	used := map[int]bool{}

	longest := 0

	for _, v := range arr {
		// 已经使用过了
		if _, ok := used[v]; ok {
			continue
		}

		length := 1
		used[v] = true

		// 元素自增查找
		for j := v + 1; ;j++ {
			if _, ok := used[j]; ok {
				used[j] = true
				length += 1
			} else {
				break
			}
		}

		// 元素自减查找
		for j := v - 1; ; j-- {
			if _, ok := used[j]; ok {
				used[j] = true
				length += 1
			} else {
				break
			}
		}

		// 更新longest值，保持最大
		if length > longest {
			longest = length
		}
	}
	return longest
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2, 9, 5}
	re := longestConsecutive(arr)
	fmt.Println(re)
}