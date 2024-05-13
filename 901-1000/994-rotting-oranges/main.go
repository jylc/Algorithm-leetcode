package main

// https://leetcode.cn/problems/rotting-oranges/description/?envType=daily-question&envId=2024-05-13
// 994. 腐烂的橘子
type point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	points := make([]point, 0)
	dir := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	m, n := len(grid), len(grid[0])
	count := 0
	ret := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				points = append(points, point{i, j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}
	for len(points) > 0 {
		length := len(points)
		for i := 0; i < length; i++ {
			p := points[i]
			for _, d := range dir {
				x, y := p.x+d[0], p.y+d[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						points = append(points, point{x, y})
						count--
						grid[x][y] = 2
					}
				}
			}
		}
		points = points[length:]
		ret++
	}
	if count != 0 {
		return -1
	}

	return ret
}
func main() {

}
