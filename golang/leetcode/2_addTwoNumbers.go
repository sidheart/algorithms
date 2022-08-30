package leetcode

/**
 * LeetCode's definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a, b, carry int
	var ans, tail *ListNode
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
			l2 = l2.Next
		}
		curVal := (a + b + carry)
		carry = curVal / 10
		curDigit := curVal % 10
		if tail == nil {
			ans = &ListNode{curDigit, nil}
			tail = ans
		} else {
			tail.Next = &ListNode{curDigit, nil}
			tail = tail.Next
		}
	}
	return ans
}
