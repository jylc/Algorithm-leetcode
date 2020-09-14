package main

//94. 二叉树的中序遍历
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
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

func inorderTraversal(root *TreeNode) []int {
	var arr []int
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node.Left)
		arr = append(arr, node.Val)
		search(node.Right)
	}
	search(root)
	return arr
}
func main() {

}
