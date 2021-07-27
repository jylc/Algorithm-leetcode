package main

//https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/
//671. 二叉树中第二小的节点

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

func findSecondMinimumValue(root *TreeNode) int {
	var dfs func(root *TreeNode)
	minVal := root.Val
	ans := -1
	dfs = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > minVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func main() {

}
