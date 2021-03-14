package main

//5702. 找出星型图的中心节点
//https://leetcode-cn.com/contest/weekly-contest-232/problems/find-center-of-star-graph/
func findCenter(edges [][]int) int {
	a, b := edges[0][0], edges[0][1]
	c, d := edges[1][0], edges[1][1]
	if a == c || a == d {
		return a
	}
	return b
}
func main() {

}
