package chapter_2

func palindrome(a *ListNode) bool {
	var reversed *ListNode
	aItr := a
	for aItr != nil {
		reversed = &ListNode{aItr.Value, reversed}
		aItr = aItr.Next
	}
	return a.Equals(reversed)
}
