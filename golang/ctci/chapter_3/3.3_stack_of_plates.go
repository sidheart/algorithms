package chapter_3

type SetOfStacks struct {
	stacks []*Stack
	capacity int
}

func NewSetOfStacks(capacity int) *SetOfStacks {
	stack := NewStack(capacity)
	return &SetOfStacks{[]*Stack{stack}, capacity}
}

func (s *SetOfStacks) IsEmpty() bool {
	if len(s.stacks) == 0 {
		return true
	}
	for _, stack := range s.stacks {
		if !stack.IsEmpty() {
			return false
		}
	}
	return true
}

func (s *SetOfStacks) Push(v interface{}) {
	n := len(s.stacks)
	if n == 0 || s.stacks[n-1].len() == s.capacity {
		s.stacks = append(s.stacks, NewStack(s.capacity))
		s.stacks[n].Push(v)
		return
	}
	s.stacks[n-1].Push(v)
}

func (s *SetOfStacks) Pop() interface{} {
	if s.IsEmpty() {
		panic("Pop() called on empty stack")
	}
	n := len(s.stacks)
	topStack := s.stacks[n-1]
	ret := topStack.Pop()
	if topStack.IsEmpty() {
		s.stacks = s.stacks[:n-1]
	}
	return ret
}

func (s *SetOfStacks) Peek() interface{} {
	if s.IsEmpty() {
		panic("Peek() called on empty stack")
	}
	n := len(s.stacks)
	topStack := s.stacks[n-1]
	return topStack.Peek()
}
