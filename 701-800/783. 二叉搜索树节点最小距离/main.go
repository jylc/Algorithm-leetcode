package main

import "math"

//783. 二叉搜索树节点最小距离
//https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	val, pre := math.MaxInt64, -1
	var midSearch func(node *TreeNode)
	midSearch = func(node *TreeNode) {
		if node == nil {
			return
		}
		midSearch(node.Left)
		if pre != -1 && node.Val-pre < val {
			val = node.Val - pre
		}
		pre = node.Val
		midSearch(node.Right)
	}
	midSearch(root)
	return val
}

func main() {

}
