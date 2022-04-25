package main

import (
	"fmt"
	"sort"
)

//interview 10.02. 变位词组
//https://leetcode-cn.com/problems/group-anagrams-lcci/
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		s := sortString(str)
		strMap[s] = append(strMap[s], str)
	}
	ans := make([][]string, 0)
	for _, v := range strMap {
		ans = append(ans, v)
	}
	fmt.Println(ans)
	return ans
}

func sortString(str string) string {
	tmp := []byte(str)
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i] < tmp[j] {
			return true
		}
		return false
	})
	return string(tmp)
}

func main() {

}
