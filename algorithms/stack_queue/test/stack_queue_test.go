package stack_queue_test

import (
	dsa "algorithms/stack_queue/src"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := new(dsa.StackQueue)

	q.Enqueue(1)
	if 1 == q.Size() {
		t.Log("queue size and enqueue succeed")
	} else {
		t.Error("queue size and enqueue failed")
	}
}

func TestDequeue(t *testing.T) {
	q := new(dsa.StackQueue)

	d1 := q.Dequeue()
	if d1 == nil {
		t.Log("empty queue dequeue succeed")
	} else {
		t.Error("empty queue dequeue failed")
	}

	q.Enqueue(1)
	d2 := q.Dequeue()
	if 1 == d2.(int) {
		t.Log("queue dequeue succeed")
	} else {
		t.Error("queue dequeue failed")
	}
}

func TestPeek(t *testing.T) {
	q := new(dsa.StackQueue)

	d1 := q.Peek()
	if d1 == nil {
		t.Log("empty queue peek succeed")
	} else {
		t.Error("empty queue peek failed")
	}

	q.Enqueue(1)
	d2 := q.Peek()
	if d2.(int) == 1 {
		t.Log("queue peek succeed")
	} else {
		t.Error("queue peek failed")
	}
}
