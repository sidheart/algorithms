package chapter_3

import "container/list"

type Queue struct {
	queue *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Enqueue(v interface{}) {
	q.queue.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Dequeue called on empty queue")
	}
	return q.queue.Remove(q.queue.Front())
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		panic("Dequeue called on empty queue")
	}
	return q.queue.Front().Value
}

func (q *Queue) IsEmpty() bool {
	return q.queue.Len() == 0
}
