package main

import (
	"fmt"
	"sort"
)

//49. 字母异位词分组
//https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := sortStr(str)
		m[s] = append(m[s], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		if v != nil {
			res = append(res, v)
		}
	}
	return res
}

func sortStr(str string) string {
	tmp := []rune(str)
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i] < tmp[j] {
			return true
		} else {
			return false
		}
	})
	return string(tmp)
}
func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(input)
	fmt.Print(res)
}
