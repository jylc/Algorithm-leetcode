package main

import "fmt"

// https://leetcode.cn/problems/minimum-length-of-string-after-deleting-similar-ends/
// 1750. 删除字符串两端相同字符后的最短长度

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	arr := []byte(s)
	for left < right {
		if arr[left] == arr[right] {
			for left+1 < right && arr[left+1] == arr[left] {
				left++
			}
			for right-1 > left && arr[right-1] == arr[right] {
				right--
			}
			left++
			right--
		} else {
			return right - left + 1
		}
	}
	return right - left + 1
}
func main() {
	arr := "bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"
	ans := minimumLength(arr)
	fmt.Println(ans)
}
