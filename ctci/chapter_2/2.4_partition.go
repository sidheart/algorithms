package chapter_2

func partition(l *ListNode, k int) *ListNode {
	dummyLeft, dummyRight := &ListNode{}, &ListNode{}
	left, right := dummyLeft, dummyRight
	for l != nil {
		v := l.Value
		if v.(int) < k {
			left.Next = &ListNode{v, nil}
			left = left.Next
		} else {
			right.Next = &ListNode{v, nil}
			right = right.Next
		}
		l = l.Next
	}
	left.Next = dummyRight.Next
	return dummyLeft.Next
}
