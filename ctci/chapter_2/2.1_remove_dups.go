package chapter_2

func removeDups(head *ListNode) {
	seen := make(map[interface{}]struct{})
	var prev *ListNode
	for head != nil {
		if _, ok := seen[head.Value]; ok {
			prev.Next = head.Next
		} else {
			seen[head.Value] = struct{}{}
			prev = head
		}
		head = head.Next
	}
}

func removeDupsNoBuffer(head *ListNode) {
	for head != nil {
		prev, runner := head, head.Next
		for runner != nil {
			if runner.Value == head.Value {
				prev.Next = runner.Next
			} else {
				prev = runner
			}
			runner = runner.Next
		}
		head = head.Next
	}
}