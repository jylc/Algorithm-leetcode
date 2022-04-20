package main

import "math"

//https://leetcode-cn.com/problems/shortest-distance-to-a-character/
//821. 字符的最短距离
func shortestToChar(s string, c byte) []int {
	var pos []int
	answer := make([]int, len(s))
	for i := range s {
		tmp := byte(s[i])
		if tmp == c {
			pos = append(pos, i)
		}
	}
	for i := range answer {
		min := math.MaxInt8
		for _, v := range pos {
			tmp := abs(i, v)
			if tmp < min {
				min = tmp
			}
		}
		answer[i] = min
	}

	return answer
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	} else {
		return a - b
	}
}
func main() {

}
