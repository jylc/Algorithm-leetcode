package main

import "sort"

//973. 最接近原点的 K 个点
//https://leetcode-cn.com/problems/k-closest-points-to-origin/

//小顶堆
func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:K]
}

func main() {

}
