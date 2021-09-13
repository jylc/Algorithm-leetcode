package main

import (
	"math"
)

//https://leetcode-cn.com/problems/number-of-boomerangs/
//447. 回旋镖的数量

//超时
func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	cnt := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			a := math.Pow(float64(points[i][0]-points[j][0]), 2) + math.Pow(float64(points[i][1]-points[j][1]), 2)
			for k := j + 1; k < n; k++ {
				b := math.Pow(float64(points[i][0]-points[k][0]), 2) + math.Pow(float64(points[i][1]-points[k][1]), 2)
				c := math.Pow(float64(points[j][0]-points[k][0]), 2) + math.Pow(float64(points[j][1]-points[k][1]), 2)
				if a == b && a == c {
					cnt += 6
				} else if a == b || a == c || b == c {
					cnt += 2
				}
			}
		}
	}
	return cnt
}

func numberOfBoomerangs1(points [][]int) int {
	ans := 0
	for _, p := range points {
		cnt := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[dis]++
		}
		for _, m := range cnt {
			ans += m * (m - 1)
		}
	}
	return ans
}

func main() {

}
