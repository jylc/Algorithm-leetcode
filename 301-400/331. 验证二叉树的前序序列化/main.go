package main

import "strings"

//331. 验证二叉树的前序序列化
//https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/

//1、栈，时间复杂度O(N)，空间复杂度O(N)
func isValidSerialization(preorder string) bool {
	stack := []int{1}
	n := len(preorder)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] != ',' {
			if preorder[i] == '#' {
				stack[len(stack)-1]--
				if stack[len(stack)-1] == 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				//读取一个数字
				for i < n && preorder[i] != ',' {
					i++
				}
				stack[len(stack)-1]--
				if stack[len(stack)-1] == 0 {
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, 2)
			}
		}
	}
	return len(stack) == 0
}

//2、栈
func isValidSerialization1(preorder string) bool {
	stack := make([]uint8, 0)
	split := strings.Split(preorder, ",")
	for _, v := range split {
		stack = append(stack, v[0])
		for len(stack) >= 3 && stack[len(stack)-1] == '#' && stack[len(stack)-2] == '#' && stack[len(stack)-3] != '#' {
			stack = stack[:len(stack)-3]
			stack = append(stack, '#')
		}
	}
	return len(stack) == 1 && stack[len(stack)-1] == '#'
}

func main() {
	isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")
}
