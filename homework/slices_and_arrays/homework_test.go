package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type AnyIntNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type AnyIntNumberSlice[T AnyIntNumber] []T

type CircularQueue[T AnyIntNumber] struct {
	values AnyIntNumberSlice[T]
	head   int
	tail   int
}

func NewCircularQueue[T AnyIntNumber](size T) CircularQueue[T] {
	return CircularQueue[T]{
		values: make(AnyIntNumberSlice[T], size),
		head:   -1,
		tail:   -1,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}
	if q.head == -1 {
		q.head = 0
	}
	q.tail = (q.tail + 1) % cap(q.values)
	q.values[q.tail] = value

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}
	if q.tail == q.head {
		q.head = -1
		q.tail = -1
		return true
	}
	q.values[q.head] = 0
	q.head = (q.head + 1) % cap(q.values)

	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.head == -1 {
		return T(q.head)
	}
	return q.values[q.head]
}

func (q *CircularQueue[T]) Back() T {
	if q.tail == -1 {
		return T(q.tail)
	}
	return q.values[q.tail]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.head == -1
}

func (q *CircularQueue[T]) Full() bool {
	return (q.head == q.tail+1) || (q.head == 0 && q.tail == cap(q.values)-1)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, []int(queue.values)))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, []int(queue.values)))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
