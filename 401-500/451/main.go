package main

import (
	"bytes"
	"fmt"
	"sort"
)

//451. 根据字符出现频率排序
//https://leetcode-cn.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	cmap := make(map[byte]int)
	for i := range s {
		cmap[s[i]]++
	}

	type pair struct {
		c   byte
		cnt int
	}
	pairs := make([]pair, 0, len(cmap))
	for k, v := range cmap {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.c}, p.cnt)...)
	}
	return string(ans)
}
func main() {
	res := frequencySort("hello")
	fmt.Println(res)
}
