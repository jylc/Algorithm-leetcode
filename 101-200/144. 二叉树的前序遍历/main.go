package main

//144. 二叉树的前序遍历
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

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

func preorderTraversal(root *TreeNode) []int {
	node := root
	stack := make([]*TreeNode, 0)
	arr := make([]int, 0)

	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			arr = append(arr, node.Val)
			node = node.Left
		} else {
			node = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}

	return arr
}
func main() {

}
