package main

//501. 二叉搜索树中的众数
//https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
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

func findMode1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if _, ok := m[node.Val]; ok {
			m[node.Val]++
		} else {
			m[node.Val] = 0
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	max := m[root.Val]
	nums := make([]int, 0)
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	for i, v := range m {
		if v == max {
			nums = append(nums, i)
		}
	}
	return nums
}

func findMode2(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}

func main() {

}
