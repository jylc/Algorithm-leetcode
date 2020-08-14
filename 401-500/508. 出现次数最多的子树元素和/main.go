package main

//508. 出现次数最多的子树元素和
//https://leetcode-cn.com/problems/most-frequent-subtree-sum/

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

func findFrequentTreeSum(root *TreeNode) []int {
	count := make(map[int]int)
	var res []int
	max := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		sum := node.Val + l + r
		if _, ok := count[sum]; ok {
			count[sum]++
		} else {
			count[sum] = 1
		}
		if max < count[sum] {
			max = count[sum]
		}
		return sum
	}
	dfs(root)
	for i, v := range count {
		if v == max {
			res = append(res, i)
		}
	}
	return res
}
func main() {

}
