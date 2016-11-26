package graph

import "testing"

func TestQueue_Enqueue_EmptyQueue(t *testing.T) {
	q := &Queue{}
	node := &Node{}

	q.Enqueue(node)

	if q.first.value != node || q.last.value != node {
		t.Error("Incorrect first and last nodes")
	}
}

func TestQueue_Enqueue_QueueWithNode(t *testing.T) {
	q := &Queue{}
	q.size = 2
	node1 := &Node{}
	node2 := &Node{}
	node3 := &Node{}
	n1 := &node{value: node1}
	n2 := &node{value: node2}
	q.first = n1
	q.first.next = n2
	q.last = n2

	q.Enqueue(node3)

	if q.first != n1 || q.first.next != n2 || n2.next.value != node3 || q.last.value != node3 || q.size != 3 {
		t.Error()
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := &Queue{}
	q.size = 2
	node1 := &Node{}
	node2 := &Node{}
	n1 := &node{value: node1}
	n2 := &node{value: node2}
	q.first = n1
	q.first.next = n2
	q.last = n2

	v, err := q.Dequeue()

	if err != nil {
		t.Error(err)
	}
	if v != node1 {
		t.Error("Incorrect node")
	}
	if q.size != 1 {
		t.Error("Incorrect size: ", q.size)
	}
}
