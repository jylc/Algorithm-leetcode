package main

// https://leetcode.cn/problems/determine-color-of-a-chessboard-square/
// 1812. 判断国际象棋棋盘中一个格子的颜色

func squareIsWhite(coordinates string) bool {
	arr := []rune(coordinates)
	x := arr[0] - 'a' + 1
	y := arr[1] - '0'
	if x%2 != 0 {
		if y%2 != 0 {
			return false
		} else {
			return true
		}
	} else {
		if y%2 != 0 {
			return true
		} else {
			return false
		}
	}
}

func main() {

}
