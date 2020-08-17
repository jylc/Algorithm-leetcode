package main

import "fmt"

//733. 图像渲染
//https://leetcode-cn.com/problems/flood-fill/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	row, col := len(image), len(image[0])
	if row == 0 || col == 0 {
		return nil
	}
	val := image[sr][sc]
	//val与newColor相同会造成dfs的死循环
	if val == newColor {
		return image
	}
	var dfs func(image [][]int, x, y int)
	dfs = func(image [][]int, x, y int) {
		if x < 0 || y < 0 || x > row-1 || y > col-1 || image[x][y] != val {
			return
		}

		image[x][y] = newColor
		dfs(image, x+1, y)
		dfs(image, x-1, y)
		dfs(image, x, y+1)
		dfs(image, x, y-1)

	}
	dfs(image, sr, sc)
	return image
}
func main() {
	image := [][]int{
		{0, 0, 0}, {0, 1, 1},
	}
	result := floodFill(image, 1, 1, 1)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}
