package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/binary-gap/
//868. 二进制间距
func binaryGap(n int) int {
	arr := make([]int, 0)
	var bin func(num int)
	bin = func(num int) {
		if num == 0 {
			return
		}
		bin(num / 2)
		arr = append(arr, num%2)
	}
	bin(n)
	fmt.Println(arr)
	i := len(arr) - 1
	flag := false
	max := 0
	tmp := 0
	for i >= 0 {
		j := i
		cnt := 0
		for j >= 0 && arr[j] != 1 {
			j--
			if flag {
				cnt++
			}
		}
		i = j
		for i >= 0 && arr[i] == 1 {
			i--
		}
		if j-i > 1 {
			tmp++
		}

		if max < cnt {
			max = cnt
		}
		if !flag {
			flag = true
		}
	}
	if max != 0 {
		max++
	} else if max < tmp {
		max = tmp
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func binaryGap1(n int) int {
	pre := -1
	cur := 0
	ans := 0
	for n != 0 {
		if n&1 == 1 {
			if pre != -1 {
				ans = max(ans, cur-pre)
			}
			pre = cur
		}
		n >>= 1
		cur++
	}
	return ans
}

func main() {
	ans := binaryGap1(22)
	fmt.Println(ans)
}
