package main

import "fmt"

//589. N 叉树的前序遍历
//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var arr []int
	var recursion func(node *Node)
	recursion = func(node *Node) {
		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		if node.Children == nil {
			return
		}

		for _, child := range node.Children {
			recursion(child)
		}
	}
	recursion(root)
	return arr
}

func preorder1(root *Node) []int {
	var iteration func(node *Node)
	var arr []int
	iteration = func(node *Node) {
		stack := make([]*Node, 0)
		if node == nil {
			return
		}
		if node.Children == nil {
			arr = append(arr, node.Val)
			return
		}

		stack = append(stack, node)
		for len(stack) != 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, curr.Val)
			for i := len(curr.Children) - 1; i >= 0; i-- {
				stack = append(stack, curr.Children[i])
			}
		}
	}
	iteration(root)
	return arr
}

func main() {
	tmp := []int{1, 2, 3, 4}

	fmt.Println(tmp[:len(tmp)-1])
	fmt.Println(tmp[len(tmp)-1])
}
