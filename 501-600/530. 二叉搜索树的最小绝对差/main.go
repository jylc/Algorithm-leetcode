package main

import (
	"math"
)

//530. 二叉搜索树的最小绝对差
//https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
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

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1

	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		f(node.Right)
	}
	f(root)

	return ans
}
func main() {

}
