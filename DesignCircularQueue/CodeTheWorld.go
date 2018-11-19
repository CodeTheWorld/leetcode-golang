package main

import "fmt"

func main() {
	queue := Constructor(3)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	fmt.Println(queue.Rear())
	fmt.Println(queue.IsFull())
	queue.DeQueue()
	queue.EnQueue(4)
	fmt.Println(queue.Rear())
}

type MyCircularQueue struct {
	queue    []int
	sizemask int
	head     int
	tail     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k+1), k + 1, 0, 0}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = value
	this.tail = this.getLogicIndex(this.tail + 1)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.getLogicIndex(this.head + 1)
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.getLogicIndex(this.tail-1)]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.head == this.getLogicIndex(this.tail+1)
}

func (this *MyCircularQueue) getLogicIndex(index int) int {
	if index < 0 {
		return index + this.sizemask
	}
	if index >= this.sizemask {
		return index - this.sizemask
	}
	return index
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
