package main

import (
	"fmt"
	"sort"
)

//275. H 指数 II
//https://leetcode-cn.com/problems/h-index-ii/
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	cnt := 0
	for i := 0; i < len(citations); i++ {
		if i+1 > citations[i] {
			break
		}
		cnt = i + 1
	}
	return cnt
}

func main() {
	index := []int{0, 1, 3, 5, 6}
	i := hIndex(index)
	fmt.Println("res", i)
}
