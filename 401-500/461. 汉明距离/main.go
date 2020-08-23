package main

//461. 汉明距离
//https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	i := x ^ y
	n := 0
	for i != 0 {
		if (i & 1) == 1 {
			n++
		}
		i = i >> 1
	}
	return n
}
func main() {

}
