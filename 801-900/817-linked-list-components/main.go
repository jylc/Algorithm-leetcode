package main

// https://leetcode.cn/problems/linked-list-components/
// 817. 链表组件
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(n)，需要遍历一遍链表。
// 空间复杂度：O(m)，其中 m 是数组 nums 的长度，需要一个哈希集合来存储 nums 的元素。

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	p := head
	count := 0
	flag := false
	for p != nil {
		v := p.Val
		if _, ok := m[v]; ok {
			if flag == false {
				flag = true
			}
		} else {
			if flag == true {
				count++
				flag = false
			}
		}
		p = p.Next
	}
	if p == nil && flag {
		count++
	}
	return count
}
func main() {

}
