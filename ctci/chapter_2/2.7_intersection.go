package chapter_2

func intersection(a, b *ListNode) bool {
	if a == nil || b == nil {
		return false
	}
	aTail, bTail := a, b
	for aTail.Next != nil {
		aTail = aTail.Next
	}
	for bTail.Next != nil {
		bTail = bTail.Next
	}
	return aTail == bTail
}

func getIntersection(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		return nil
	}
	aTail, bTail := a, b
	m, n := 1, 1
	for aTail.Next != nil {
		aTail = aTail.Next
		m++
	}
	for bTail.Next != nil {
		bTail = bTail.Next
		n++
	}
	if aTail != bTail {
		return nil
	}
	aItr, bItr := a, b
	if m > n {
		for i := 0; i < m - n; i++ {
			aItr = aItr.Next
		}
	}
	if n > m {
		for i := 0; i < m - n; i++ {
			bItr = bItr.Next
		}
	}
	for aItr != bItr {
		aItr, bItr = aItr.Next, bItr.Next
	}
	return aItr
}
