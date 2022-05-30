package main

//https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
//1022. 从根到叶的二进制数之和
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

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, val int) int
	dfs = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}
		val = val<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			return val
		}
		return dfs(node.Left, val) + dfs(node.Right, val)
	}

	return dfs(root, 0)
}

func main() {
}
