package utils

import (
	"github.com/rotisserie/eris"
)

type Queue[T any] struct {
	data    []T
	pointer int
}

func NewQueue[T any](items ...T) *Queue[T] {
	return &Queue[T]{
		data:    append([]T{}, items...),
		pointer: 0,
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(data ...T) {
	q.data = append(q.data, data...)
}

// Dequeue removes and returns the element at the front of the queue
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, eris.New("Queue is empty")
	}

	val := q.data[q.pointer]
	q.pointer++

	if q.pointer > 100 && q.pointer*2 >= len(q.data) {
		q.data = append([]T{}, q.data[q.pointer:]...)
		q.pointer = 0
	}

	return val, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.pointer >= len(q.data)
}
