package main

// https://leetcode.cn/problems/max-chunks-to-make-sorted/
// 769. 最多能完成排序的块

func maxChunksToSorted(arr []int) int {
	max := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			ans++
		}
	}
	return ans
}
func main() {

}
