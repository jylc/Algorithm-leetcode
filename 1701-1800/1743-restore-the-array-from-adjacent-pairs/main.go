package main

//https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs/
//1743. 从相邻元素对还原数组

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	numsMap := make(map[int][]int, n)
	for _, p := range adjacentPairs {
		v, w := p[0], p[1]
		numsMap[v] = append(numsMap[v], w)
		numsMap[w] = append(numsMap[w], v)
	}
	minLen := 1
	idx := 0
	visited := make(map[int]bool)
	for k, v := range numsMap {
		if len(v) == minLen {
			minLen = len(v)
			idx = k
		}
		visited[k] = false
	}
	ans := make([]int, 0)
	var dfs func(m map[int][]int)
	dfs = func(m map[int][]int) {
		if len(ans) == len(m) {
			return
		}
		nums := m[idx]
		ans = append(ans, idx)
		visited[idx] = true
		for _, num := range nums {
			if visited[num] == false {
				idx = num
				dfs(m)
			}
		}
	}
	dfs(numsMap)
	return ans
}

func main() {
	arr := [][]int{{1, 2}, {3, 4}, {3, 5}, {2, 5}}
	restoreArray(arr)
}
