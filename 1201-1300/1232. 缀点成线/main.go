package main

import "fmt"

//1232. 缀点成线
//https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/
func checkStraightLine(coordinates [][]int) bool {
	row, _ := len(coordinates), len(coordinates[0])
	if row == 2 {
		return true
	}
	a, b := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]
	for i := 2; i < row; i++ {
		x2, y2 := coordinates[i][0], coordinates[i][1]
		//将除法((y1-b)/(x1-a))==((y2-b)/(x2-a))转为乘法((y1-b)*(x2-a))==((y2-b)*(x1-a))，避免出现乘0
		if float64(y1-b)*float64(x2-a) != float64(y2-b)*float64(x1-a) {
			return false
		}
	}
	return true
}
func main() {
	arr := [][]int{
		{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7},
	}
	result := checkStraightLine(arr)
	fmt.Println("result = ", result)
}
