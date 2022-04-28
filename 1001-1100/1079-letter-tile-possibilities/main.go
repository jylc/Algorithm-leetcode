package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/letter-tile-possibilities/
//1079. 活字印刷
func numTilePossibilities(tiles string) int {
	visited := make([]bool, len(tiles))
	bytes := []byte(tiles)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	res := 0
	var backTrace func(point int)
	backTrace = func(point int) {
		if point == len(tiles) {
			return
		}

		for i := 0; i < len(tiles); i++ {
			if visited[i] || (i > 0 && !visited[i-1] && bytes[i] == bytes[i-1]) {
				continue
			}
			res++
			visited[i] = true
			backTrace(point + 1)
			visited[i] = false
		}
	}
	backTrace(0)
	return res
}

func numTilePossibilities1(tiles string) int {
	n := len(tiles)
	//map用来消除重复字母
	m := make(map[uint8]int)
	ans := 0
	for i := 0; i < n; i++ {
		m[tiles[i]]++
	}
	var dfs func() int
	dfs = func() int {
		res := 0
		//不重复地选择一个字母
		for k, v := range m {
			if v == 0 {
				continue
			}
			res++
			m[k]--
			res += dfs()
			m[k]++
		}
		return res
	}
	ans = dfs()
	return ans
}
func main() {
	ans := numTilePossibilities1("AAB")
	fmt.Println(ans)
}
