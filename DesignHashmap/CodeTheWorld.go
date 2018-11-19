package main

import "fmt"

func main() {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	hashMap.Put(10002, 2)
	fmt.Println(hashMap.Get(1))
	fmt.Println(hashMap.Get(3))
	hashMap.Put(2, 1)
	fmt.Println(hashMap.Get(2))
	hashMap.Remove(2)
	fmt.Println(hashMap.Get(2))
}

/**
思路：拉链法实现hash表
*/
type KVNode struct {
	Key  int
	Val  int
	Next *KVNode
}

type MyHashMap struct {
	hashMap  []*KVNode
	sizemask int
}

func Constructor() MyHashMap {
	return MyHashMap{make([]*KVNode, 10000), 10000}
}

func (this *MyHashMap) Put(key int, value int) {
	node := this.hashMap[this.hash(key)]
	ptr := node
	for ; ptr != nil; ptr = ptr.Next {
		if key == ptr.Key {
			ptr.Val = value
			break
		}
	}
	if nil == ptr {
		this.hashMap[this.hash(key)] = &KVNode{key, value, node}
	}
}

func (this *MyHashMap) Get(key int) int {
	node := this.hashMap[this.hash(key)]
	ptr := node
	for ; nil != ptr; ptr = ptr.Next {
		if ptr.Key == key {
			return ptr.Val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	node := this.hashMap[this.hash(key)]
	ptr := node
	if nil == node {
		return
	}
	if node.Key == key {
		this.hashMap[this.hash(key)] = node.Next
		return
	}
	for ; nil != ptr.Next; ptr = ptr.Next {
		if ptr.Next.Key == key {
			ptr.Next = ptr.Next.Next
			break
		}
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.sizemask
}
