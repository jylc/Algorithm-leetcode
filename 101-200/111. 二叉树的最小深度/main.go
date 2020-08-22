package main

//111. 二叉树的最小深度
//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/submissions/
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	//添加这条语句是防止一个父节点只有一个子节点时返回节点数0
	//例如[1 2]，父节点1有子节点2，最小深度不为1，而是2
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}
func main() {

}
