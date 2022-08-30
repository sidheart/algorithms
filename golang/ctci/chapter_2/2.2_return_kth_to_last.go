package chapter_2

func returnKthToLast(l *ListNode, k int) *ListNode {
	if l == nil {
		return nil
	}
	n := l.Len()
	advance := n - 1 - k
	if advance < 0 {
		return nil
	}
	for i := 0; i < advance; i++ {
		l = l.Next
	}
	return l
}
