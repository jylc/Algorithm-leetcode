package main

//145. 二叉树的后序遍历
//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
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

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	node := root
	stack := make([]*TreeNode, 0)
	arr := make([]int, 0)
	for node != nil || len(stack) != 0 {
		if node != nil {
			arr = append(arr, node.Val)
			stack = append(stack, node)
			node = node.Right
		} else {
			length := len(stack)
			node = stack[length-1].Left
			stack = stack[:length-1]
		}
	}
	length := len(arr)
	for i := 0; i < length/2; i++ {
		tmp := arr[length-i-1]
		arr[length-i-1] = arr[i]
		arr[i] = tmp
	}
	return arr
}
func main() {

}
