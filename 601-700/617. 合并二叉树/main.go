package main

//617. 合并二叉树
//https://leetcode-cn.com/problems/merge-two-binary-trees/
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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	root := &TreeNode{Val: t1.Val + t2.Val}
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}
func main() {

}
