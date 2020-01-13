package chapter_3

type SortStack struct {
	stack *MinStack
}

func NewSortStack(initialCapacity int) *SortStack {
	return &SortStack{NewMinStack(initialCapacity)}
}

func (s *SortStack) IsEmpty() bool {
	return s.stack.IsEmpty()
}

func (s *SortStack) Push(v int) {
	s.stack.Push(v)
}

func (s *SortStack) sort() {
	tempStack := NewMinStack(s.stack.Size())
	for !s.stack.IsEmpty() {
		e := s.stack.Pop()
		if tempStack.IsEmpty() {
			tempStack.Push(e)
			continue
		}
		count := 0
		for !tempStack.IsEmpty() && tempStack.Peek() < e {
			s.stack.Push(tempStack.Pop())
			count++
		}
		tempStack.Push(e)
		for i := 0; i < count; i++ {
			tempStack.Push(s.stack.Pop())
		}
	}
	s.stack = tempStack
}

func (s *SortStack) Peek() int {
	if s.IsEmpty() {
		panic("Peek() called on empty stack")
	}
	top := s.stack.Peek()
	if s.stack.Min() == top {
		return top
	}
	s.sort()
	return s.stack.Peek()
}

func (s *SortStack) Pop() int {
	if s.IsEmpty() {
		panic("Pop() called on empty stack")
	}
	top := s.stack.Peek()
	if s.stack.Min() == top  {
		return s.stack.Pop()
	}
	s.sort()
	return s.stack.Pop()
}



