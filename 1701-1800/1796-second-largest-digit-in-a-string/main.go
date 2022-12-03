package main

// https://leetcode.cn/problems/second-largest-digit-in-a-string/description/
// 1796. 字符串中第二大的数字

func secondHighest(s string) int {
	arr := make([]int, 10)
	for _, i := range s {
		if i >= '0' && i <= '9' {
			arr[i-'0']++
		}
	}
	count := 0
	for i := 9; i >= 0; i-- {
		if arr[i] != 0 {
			count++
			if count == 2 {
				return i
			}
		}
	}
	return -1
}

func main() {

}
