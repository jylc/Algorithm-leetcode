package main

import "fmt"

//interview 08.06. 汉诺塔问题
//https://leetcode-cn.com/problems/hanota-lcci/

func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}

	var move func(int, *[]int, *[]int, *[]int)
	move = func(num int, from *[]int, mid *[]int, to *[]int) {
		if num < 0 {
			return
		}
		if num == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}
		move(num-1, from, to, mid)
		move(1, from, mid, to)
		move(num-1, mid, from, to)
	}
	move(len(A), &A, &B, &C)
	return C
}

func main() {
	A, B, C := []int{2, 1, 0}, []int{}, []int{}
	C = hanota(A, B, C)
	fmt.Println("A", A)
	fmt.Println("B", B)
	fmt.Println("C", C)
}
