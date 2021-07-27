package main

//https://leetcode-cn.com/problems/check-if-all-the-integers-in-a-range-are-covered/
//1893. 检查是否区域内所有整数都被覆盖
func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		flag := false
		for _, r := range ranges {
			if r[0] <= i && r[1] >= i {
				flag = true
				continue
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}
func main() {

}
