package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(9)
	dq := q.Dequeue()
	n := q.Len()
	if q.Len() != 3 {
		t.Fatal("Queue length should be equal to 2, got %d", n)
	}
	if dq != 2 {
		t.Fatalf("Dequeued value should be equal to 2, got %d", dq)
	}
}
