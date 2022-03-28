package main

//https://leetcode-cn.com/problems/kplEvH/
//LCP 45. 自行车炫技赛场
var offset = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) (ans [][]int) {
	n, m := len(terrain), len(terrain[0])
	res := make([][][102]bool, n)
	for i := range res {
		res[i] = make([][102]bool, m)
	}

	x, y := position[0], position[1]
	res[x][y][1] = true
	type pair struct {
		x, y, speed int
	}

	q := []pair{{x, y, 1}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, speed := p.x, p.y, p.speed
		for _, d := range offset {
			if xx, yy := x+d.x, y+d.y; xx >= 0 && yy >= 0 && xx < n && yy < m {
				s := speed + terrain[x][y] - terrain[xx][yy] - obstacle[xx][yy]
				if s > 0 && !res[xx][yy][s] {
					res[xx][yy][s] = true
					q = append(q, pair{xx, yy, s})
				}
			}
		}
	}

	res[x][y][1] = false
	for i, row := range res {
		for j, b := range row {
			if b[1] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}

func main() {

}
