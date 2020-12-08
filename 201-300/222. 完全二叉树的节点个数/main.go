package main

//222. 完全二叉树的节点个数
//https://leetcode-cn.com/problems/count-complete-tree-nodes/
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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	left = countNodes(root.Left) + 1
	right = countNodes(root.Right) + 1
	return left + right
}
func main() {

}
