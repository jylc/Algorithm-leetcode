package main

import (
	"fmt"
	"sort"
)

//332. 重新安排行程
//https://leetcode-cn.com/problems/reconstruct-itinerary/
//与图相关
func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)
	for _, ticket := range tickets {
		src, des := ticket[0], ticket[1]
		m[src] = append(m[src], des)
	}
	for key := range m {
		sort.Strings(m[key])
	}
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
func main() {
	arr := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	result := findItinerary(arr)
	fmt.Println("result = ", result)
}
