package playgo

import (
	"fmt"
	"sync"
)

// BBQueue implements a queue which will block when there a caller attempts to push a new element
// and the maximum capacity of the queue has been reached
type BBQueue interface {
	Enqueue(element int)
	Dequeue() int
	Size() int
}

// BoundedBlockingQueue stucture used to implement a queue of limited capacity which will block
// upon Enqueue invocation if the capacity is reached.
type BoundedBlockingQueue struct {
	Capacity int
	size     int
	q        []int
	head     int
	tail     int
	c        *sync.Cond
	m        sync.Locker
}

// NewBoundedBlockingQueue implements a constructor for the BBQ.
func NewBoundedBlockingQueue(capacity int) *BoundedBlockingQueue {

	bbq := &BoundedBlockingQueue{
		Capacity: capacity,
		size:     0,
		head:     0,
		tail:     0,
	}
	bbq.q = make([]int, capacity)
	bbq.m = new(sync.Mutex)
	bbq.c = sync.NewCond(bbq.m)
	return bbq
}

// Enqueue implements the push part of the BBQ interface. If the queue is full this call blocks
// until Dequeue is called.
func (q *BoundedBlockingQueue) Enqueue(element int) {
	// defer release the mux
	defer q.m.Unlock()
	q.m.Lock()

	// if capacity reached, block
	if q.size > q.Capacity {
		panic(fmt.Sprintf("Capacity exceded! Size: %d Capacity: %d", q.size, q.Capacity))
	}

	// Wait for someone to Dequeue
	for q.size == q.Capacity {
		q.c.Wait()
	}

	q.q[q.head] = element
	q.head = q.head + 1
	if q.head == q.Capacity {
		q.head = 0
	}
	q.size = q.size + 1
	q.c.Broadcast()
}

// Dequeue implements the pop part of the BBQ interface.
func (q *BoundedBlockingQueue) Dequeue() int {
	defer q.m.Unlock()
	q.m.Lock()

	for q.size == 0 {
		q.c.Wait()
	}

	element := q.q[q.tail]
	q.tail = q.tail + 1
	if q.tail == q.Capacity {
		q.tail = 0
	}
	q.size = q.size - 1
	q.c.Broadcast()
	return element
}

// Size returns the number of elements waiting in the BBQ.
func (q *BoundedBlockingQueue) Size() int {
	return q.size
}
