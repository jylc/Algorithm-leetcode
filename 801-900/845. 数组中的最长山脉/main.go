package main

//845. 数组中的最长山脉
//https://leetcode-cn.com/problems/longest-mountain-in-array/
func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	left := make([]int, len(A))
	right := make([]int, len(A))

	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i+1] < A[i] {
			right[i] = right[i+1] + 1
		}
	}

	ans := 0
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = 1 + r + l
		}
	}
	return ans
}

func main() {

}
