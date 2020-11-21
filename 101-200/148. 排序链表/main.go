package main

import "sort"

//148. 排序链表
//https://leetcode-cn.com/problems/sort-list/
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

func sortList1(head *ListNode) *ListNode {
	//将链表转为数组，对数组排序，再赋值到原链表中,排序时间复杂度O(nlogn)，空间复杂度O(n)
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	tmp := head
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}

	sort.Ints(arr)
	i := 0
	tmp = head
	for tmp != nil {
		tmp.Val = arr[i]
		i++
		tmp = tmp.Next
	}
	return head
}

func sortList2(head *ListNode) *ListNode {
	//自顶向下归并排序，时间复杂度O(nlogn)，空间复杂度O(logn)
	return sortL(head, nil)
}

func sortL(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	fast, slow := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sortL(head, mid), sortL(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}

func main() {

}
