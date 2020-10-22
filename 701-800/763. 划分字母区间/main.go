package main

import "fmt"

//763. 划分字母区间
//https://leetcode-cn.com/problems/partition-labels/
func partitionLabels(S string) []int {
	m := make(map[rune]bool)
	res := make([]int, 0)
	max, j := 0, 0
	t := 0
	for i, s := range S {
		for j = i; j < len(S); j++ {
			if int32(S[j]) == s {
				m[s] = true
				if max < j {
					max = j
				}
			}
		}
		if _, ok := m[s]; ok {
			if max == i {
				res = append(res, max-t+1)
				t = max + 1
			}
		}
	}
	return res
}
func main() {
	res := partitionLabels("abd")
	for _, re := range res {
		fmt.Print(re)
	}
}
