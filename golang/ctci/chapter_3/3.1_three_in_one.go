package chapter_3

import (
	"errors"
)

type StackInfo struct {
	start, end, top int
}

type MultiStack struct {
	numStacks int
	stack []interface{}
	bounds []StackInfo
}

func NewMultiStack(numStacks, initialCapacity int) (*MultiStack, error) {
	if initialCapacity < 1 {
		return nil, errors.New("initialCapacity must be >= 1")
	}
	stack := make([]interface{}, initialCapacity * numStacks)
	bounds := make([]StackInfo, 0, numStacks)
	for count, start := 0, 0; count < numStacks; count++ {
		end := start + initialCapacity
		bounds = append(bounds, StackInfo{start, end, -1})
		start = end
	}
	return &MultiStack{numStacks, stack, bounds}, nil
}

func NewThreeInOneStack(initialCapacity int) *MultiStack {
	stack, _ := NewMultiStack(3, initialCapacity)
	return stack
}

func (ms *MultiStack) isFull(stackNum int) bool {
	info := ms.bounds[stackNum]
	return info.end == info.top + 1
}

func (ms *MultiStack) growStack(stackNum int) {
	info := &ms.bounds[stackNum]
	segLen := info.end - info.start
	newStack := make([]interface{}, len(ms.stack) + segLen + 1)
	copy(newStack[:info.start], ms.stack[:info.start])
	copy(newStack[info.start:info.end], ms.stack[info.start:info.end])
	info.end += segLen + 1
	start := info.end
	for i := stackNum + 1; i < ms.numStacks; i++ {
		stackInfo := &ms.bounds[i]
		segLen = stackInfo.end - stackInfo.start
		copy(newStack[start:start+segLen], ms.stack[stackInfo.start:stackInfo.end])
		topDelta := stackInfo.top - stackInfo.start
		stackInfo.start = start
		if stackInfo.top != -1 {
			stackInfo.top = start + topDelta
		}
		stackInfo.end = start + segLen
		start = start + segLen
	}
	ms.stack = newStack
}

func (ms *MultiStack) Push(stackNum int, v interface{}) error {
	if stackNum < 0 || stackNum >= ms.numStacks {
		return errors.New("invalid stackNum specified")
	}
	if ms.isFull(stackNum) {
		ms.growStack(stackNum)
	}
	info := &ms.bounds[stackNum]
	if info.top == -1 {
		info.top = info.start
	} else {
		info.top++
	}
	ms.stack[info.top] = v
	return nil
}

func (ms *MultiStack) Pop(stackNum int) error {
	if stackNum < 0 || stackNum >= ms.numStacks {
		return errors.New("invalid stackNum specified")
	}
	info := &ms.bounds[stackNum]
	if info.top == -1 {
		return errors.New("attempted to pop from empty stack")
	}
	if info.top == info.start {
		info.top = -1
	} else {
		info.top--
	}
	return nil
}

func (ms *MultiStack) Peek(stackNum int) (interface{}, error) {
	if stackNum < 0 || stackNum >= ms.numStacks {
		return nil, errors.New("invalid stackNum specified")
	}
	info := ms.bounds[stackNum]
	if info.top == -1 {
		return nil, errors.New("attempted to peek on empty stack")
	}
	return ms.stack[info.top], nil
}

func (ms *MultiStack) IsEmpty(stackNum int) (bool, error) {
	if stackNum < 0 || stackNum >= ms.numStacks {
		return false, errors.New("invalid stackNum specified")
	}
	info := ms.bounds[stackNum]
	return info.top == -1, nil
}