package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if 0 == len(this.minStack) || x <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
