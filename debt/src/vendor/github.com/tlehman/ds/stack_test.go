package ds

import "testing"

func TestStack(t *testing.T) {
	q := NewStack()
	q.Push(2)
	q.Push(3)
	q.Push(5)
	q.Push(9)
	p := q.Pop()
	n := q.Len()
	if q.Len() != 3 {
		t.Fatal("Stack length should be equal to 2, got %d", n)
	}
	if p != 9 {
		t.Fatalf("Popped value should be equal to 2, got %d", p)
	}
}
