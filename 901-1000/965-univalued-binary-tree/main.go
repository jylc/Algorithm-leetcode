package main

//https://leetcode.cn/problems/univalued-binary-tree/
//965. 单值二叉树
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

func isUnivalTree(root *TreeNode) (ans bool) {
	var dfs func(node *TreeNode)
	value := root.Val
	ans = true
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val != value {
			ans = false
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return
}
func main() {

}
