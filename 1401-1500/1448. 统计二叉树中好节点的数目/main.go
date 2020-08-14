package main

import "fmt"

//1448. 统计二叉树中好节点的数目
//https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/

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

var (
	count = 0
)
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := root.Val
	count=0
	countIt(root,max)
	return count
}

func countIt(root *TreeNode, max int) {
	if root.Val >= max {
		max = root.Val
		fmt.Println(root.Val)
		count++
	}
	tmp := max
	if root.Left != nil {
		countIt(root.Left, tmp)
	}
	if root.Right != nil {
		tmp = max
		countIt(root.Right, tmp)
	}
}
func main() {

}
