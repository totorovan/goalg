package graph

import (
	"errors"
)

type Queue struct {
	first *node
	last  *node
	size  int
}

type node struct {
	value *Node
	next  *node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Enqueue(n *Node) {
	oldLast := q.last
	q.last = &node{value: n}
	if q.Len() == 0 {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.size++
}

func (q *Queue) Dequeue() (*Node, error) {
	if q.size == 0 {
		return nil, errors.New("Empty")
	}
	v := q.first.value
	newFirst := q.first.next
	q.first = newFirst
	q.size--
	return v, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
