package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (53.74%)
 * Total Accepted:    59.4K
 * Total Submissions: 110.6K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// if l1 == nil {
	// 	return l2
	// }
	// if l2 == nil {
	// 	return l1
	// }

	// var res *ListNode
	// if l1.Val >= l2.Val {
	// 	res = l2
	// 	res.Next = mergeTwoLists(l1, l2.Next)
	// } else {
	// 	res = l1
	// 	res.Next = mergeTwoLists(l1.Next, l2)
	// }
	// return res


	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result *ListNode = new(ListNode)

	currentNode := result

	var l1Node *ListNode = l1

	var l2Node *ListNode = l2

	for l1Node != nil || l2Node != nil {

		// 第二个链表结束
		for l1Node != nil && l2Node == nil {
			currentNode.Next = l1Node
			l1Node = l1Node.Next
			currentNode = currentNode.Next
		}
		for l1Node == nil && l2Node != nil {
			currentNode.Next = l2Node
			l2Node = l2Node.Next
			currentNode = currentNode.Next
		}

		// p1>p2
		for l1Node != nil && l2Node != nil && l1Node.Val >= l2Node.Val {
			currentNode.Next = l2Node
			l2Node = l2Node.Next
			currentNode = currentNode.Next
		}

		// p2>p1
		for l1Node != nil && l2Node != nil && l1Node.Val < l2Node.Val {
			currentNode.Next = l1Node
			l1Node = l1Node.Next
			currentNode = currentNode.Next
		}

	}

	return result.Next

}
