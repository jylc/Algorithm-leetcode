package main

//896. 单调数列
//https://leetcode-cn.com/problems/monotonic-array/

func isMonotonic(A []int) bool {
	lens := len(A)
	dec, inc := true, true

	for i := 1; i < lens; i++ {
		if A[i] < A[i-1] {
			inc = false
		}
		if A[i] > A[i-1] {
			dec = false
		}
	}
	return dec || inc
}
func main() {

}
