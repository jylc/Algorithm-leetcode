package main

import "fmt"

//https://leetcode-cn.com/problems/count-lattice-points-inside-a-circle/
//6042. 统计圆内格点数目
type point struct {
	x, y int
}

func isInCircle(p point, r int) bool {
	if p.x*p.x+p.y*p.y <= r*r {
		return true
	}
	return false
}
func countLatticePoints(circles [][]int) int {

	points := make(map[point]bool)
	index := 0
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		for a := x - r; a <= x+r; a++ {
			for b := y - r; b <= y+r; b++ {
				q := point{
					x: a,
					y: b,
				}
				p := point{
					x: a - x,
					y: b - y,
				}
				if _, ok := points[q]; !ok {
					if isInCircle(p, r) {
						points[q] = true
						index++
					}
				}
			}
		}
	}
	return index
}
func main() {
	//circles := [][]int{{2, 2, 2}, {3, 4, 1}}
	circles := [][]int{{2, 2, 1}}
	ans := countLatticePoints(circles)
	fmt.Println(ans)
}
