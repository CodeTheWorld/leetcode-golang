package main

import "fmt"

func main() {
	set := Constructor()
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(3))
	set.Add(2)
	fmt.Println(set.Contains(2))
	set.Remove(2)
	fmt.Println(set.Contains(2))
}

type MyHashSet struct {
	ht []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]int, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.ht[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.ht[key] = 0
}

/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return 1 == this.ht[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
