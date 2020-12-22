package main

/*
  给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

  示例 1:
    输入: 1->2->3->3->4->4->5
    输出: 1->2->5

  示例 2:
    输入: 1->1->1->2->3
    输出: 2->3

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	ptr := headPre
	node := head
	for node != nil {
		val := node.Val
		count := 1
		next := node.Next
		for next != nil && next.Val == val {
			count++
			next = next.Next
		}
		if count == 1 {
			ptr.Next = node
			ptr = node
		}
		node = next
	}
	ptr.Next = nil
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}