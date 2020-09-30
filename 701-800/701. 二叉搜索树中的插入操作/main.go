package main

import "fmt"

//701. 二叉搜索树中的插入操作
//https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var dfs func(node *TreeNode, val int) *TreeNode
	dfs = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}
		if node.Val > val {
			left := dfs(node.Left, val)
			if left != nil {
				node.Left = left
				left = nil
			}
		} else if node.Val < val {
			right := dfs(node.Right, val)
			if right != nil {
				node.Right = right
				right = nil
			}
		}
		return nil
	}
	dfs(root, val)
	return root
}
func main() {
	root := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	bst := insertIntoBST(root, 5)
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		fmt.Print("\n ", root.Val)
		f(root.Left)
		f(root.Right)
	}
	f(bst)
}
