package main

import "fmt"

//面试题3：二维数组中的查找
func Find(matrix []int, cols, rows, number int) bool {
	found := false
	if rows == 0 || cols == 0 || len(matrix) == 0 {
		return found
	}
	row, col := 0, cols-1

	for row < rows && col >= 0 {
		if matrix[row*cols+col] == number {
			found = true
			break
		} else if matrix[row*cols+col] > number {
			col--
		} else {
			row++
		}
	}
	return found
}

func main() {
	arr := []int{1, 2, 8, 9, 2, 4, 9, 12, 4, 7, 10, 13, 6, 8, 11, 15}
	found := Find(arr, 4, 4, 7)
	fmt.Print("result:", found)
}
