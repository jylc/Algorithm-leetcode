package main

//https://leetcode-cn.com/problems/count-unhappy-friends/
//1583. 统计不开心的朋友
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	matrix := make([][]int, n)
	for i, preference := range preferences {
		matrix[i] = make([]int, n)
		for j, p := range preference {
			matrix[i][p] = j
		}
	}

	match := make([]int, n)
	for _, pair := range pairs {
		match[pair[0]] = pair[1]
		match[pair[1]] = pair[0]
	}
	cnt := 0
	for x, y := range match {
		index := matrix[x][y]
		for _, u := range preferences[x][:index] {
			v := match[u]
			if matrix[u][x] < matrix[u][v] {
				cnt++
				break
			}
		}
	}
	return cnt
}

func main() {

}
