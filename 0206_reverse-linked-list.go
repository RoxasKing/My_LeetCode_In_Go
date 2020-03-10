package leetcode

/*
  反转一个单链表。
  示例:
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL
  进阶:
  你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var rev *ListNode
	for head != nil {
		tmp := head.Next // 当前头节点的后继节点指针
		head.Next = rev  // 当前头节点插入到反转链表首部
		rev = head       // 反转链表头部指针指向当前遍历到的头节点
		head = tmp       // 后继节点指针变为当前头节点指针
	}
	return rev
}