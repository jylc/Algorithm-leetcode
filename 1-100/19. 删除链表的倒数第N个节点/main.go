package main

//19. 删除链表的倒数第N个节点
//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, second := head, head

	for second.Next != nil {
		if n > 0 {
			n--
		} else {
			first = first.Next

		}
		second = second.Next
	}
	if first == second {
		return nil
	}
	if n > 0 {
		return head.Next
	}
	first = first.Next.Next
	return head
}

func main() {

}
