package main

import (
	"fmt"
	"math"
)

//842. 将数组拆分成斐波那契序列
//https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/
func splitIntoFibonacci(S string) []int {
	res := make([]int, 0)
	backtrace(S, &res, 0)
	return res
}

func backtrace(S string, res *[]int, index int) bool {
	if index == len(S) && len(*res) >= 3 {
		return true
	}
	for i := index; i < len(S); i++ {
		if S[index] == '0' && i > index {
			break
		}

		tmpNum := subNum(S, index, i+1)
		if tmpNum > math.MaxInt32 {
			break
		}
		size := len(*res)
		if size >= 2 && tmpNum > (*res)[size-1]+(*res)[size-2] {
			break
		}

		if size <= 1 || tmpNum == (*res)[size-1]+(*res)[size-2] {
			*res = append(*res, tmpNum)
			if backtrace(S, res, i+1) {
				return true
			}
			*res = (*res)[:len(*res)-1]
		}
	}
	return false
}
func subNum(S string, index int, end int) int {
	res := 0
	for i := index; i < end; i++ {
		res = 10*res + int(S[i]-'0')
	}
	return res
}

func main() {
	res := splitIntoFibonacci("123456579")
	fmt.Print("result = ", res)
}
