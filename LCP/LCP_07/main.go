package main

//LCP 07. 传递信息
//https://leetcode-cn.com/problems/chuan-di-xin-xi/
func numWays(n int, relation [][]int, k int) int {
	edges := make([][]int, n)
	for _, ints := range relation {
		a, b := ints[0], ints[1]
		edges[a] = append(edges[a], b)
	}

	ans := 0
	var dfs func(int, int)
	dfs = func(x int, step int) {
		if step == k {
			if x == n-1 {
				ans++
			}
			return
		}
		for _, y := range edges[x] {
			dfs(y, step+1)
		}
	}
	dfs(0, 0)
	return ans
}

func main() {
	n := 5
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	k := 3
	numWays(n, relation, k)
}
