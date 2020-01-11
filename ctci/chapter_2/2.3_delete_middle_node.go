package chapter_2

func deleteMiddleNode(m *ListNode) {
	m.Value = m.Next.Value
	m.Next = m.Next.Next
}
