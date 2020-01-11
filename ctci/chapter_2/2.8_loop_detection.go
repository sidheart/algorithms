package chapter_2

func detectLoop(a *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	aItr := a
	for aItr != nil {
		if _, ok := seen[aItr]; ok {
			return aItr
		}
		seen[aItr] = struct{}{}
		aItr = aItr.Next
	}
	return nil
}
