package main

import "fmt"

func main() {
	duQue := Constructor(3)
	duQue.InsertLast(1)
	duQue.InsertLast(2)
	duQue.InsertFront(3)
	duQue.InsertFront(4)
	fmt.Println(duQue.GetRear())
	fmt.Println(duQue.IsFull())
	duQue.DeleteLast()
	duQue.InsertFront(4)
	fmt.Println(duQue.GetFront())
}

type MyCircularDeque struct {
	queue    []int
	head     int
	tail     int
	sizemask int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k+1), 0, 0, k + 1}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = this.convertIndex(this.head - 1)
	this.queue[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = value
	this.tail = this.convertIndex(this.tail + 1)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.convertIndex(this.head + 1)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = this.convertIndex(this.tail - 1)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.convertIndex(this.tail-1)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return this.head == this.convertIndex(this.tail+1)
}

func (this *MyCircularDeque) convertIndex(index int) int {
	if index < 0 {
		return index + this.sizemask
	}
	return index % this.sizemask
}
