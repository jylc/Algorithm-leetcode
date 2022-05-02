package main

import (
	"fmt"
	"sort"
)

// TreeNode https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
//1305. 两棵二叉搜索树中的所有元素
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := make([]int, 0), make([]int, 0)
	var search func(root *TreeNode, arr []int) []int
	search = func(root *TreeNode, arr []int) []int {
		if root == nil {
			return arr
		}
		arr = search(root.Left, arr)
		arr = append(arr, root.Val)
		arr = search(root.Right, arr)
		return arr
	}

	arr1 = search(root1, arr1)
	arr2 = search(root2, arr2)
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	return arr1
}

func test(arr []int) {
	arr = append(arr, 100)
}
func main() {
	arr := make([]int, 0)
	test(arr)
	fmt.Println(arr)
}
