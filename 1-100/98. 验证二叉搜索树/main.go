package main

import "math"

//98. 验证二叉搜索树
//https://leetcode-cn.com/problems/validate-binary-search-tree/

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

func isValidBST(root *TreeNode) bool {
	val := math.MinInt64
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if left := dfs(node.Left); !left {
			return false
		}
		if node.Val <= val {
			return false
		}
		val = node.Val
		return dfs(node.Right)
	}

	return dfs(root)
}
func main() {

}
