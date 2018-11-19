package main

import "fmt"

func main() {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
	cache.Put(5, 5)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(5))
}

type LruNode struct {
	Key  int
	Val  int
	Pre  *LruNode
	Next *LruNode
	Box  *LfuNode
}

type LfuNode struct {
	Freq int
	Pre  *LfuNode
	Next *LfuNode
	Head *LruNode
	Tail *LruNode
}

type LFUCache struct {
	capacity int
	head     *LfuNode
	tail     *LfuNode
	hashMap  map[int]*LruNode
	size     int
}

func Constructor(capacity int) LFUCache {
	head := &LfuNode{-1, nil, nil, nil, nil}
	tail := &LfuNode{-1, nil, nil, nil, nil}
	head.Next = tail
	tail.Pre = head
	return LFUCache{capacity, head, tail, make(map[int]*LruNode), 0}
}

func (this *LFUCache) Get(key int) int {
	lruNode, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	this.removeNodeFromList(lruNode)
	this.insertNodeIntoBoxAfterLfuNode(lruNode, lruNode.Box.Pre, lruNode.Box.Freq+1)
	return lruNode.Val
}

func (this *LFUCache) Put(key int, value int) {
	if 0 == this.capacity {
		return
	}
	lruNode, ok := this.hashMap[key]
	if ok {
		lruNode.Val = value
		this.removeNodeFromList(lruNode)
		this.insertNodeIntoBoxAfterLfuNode(lruNode, lruNode.Box.Pre, lruNode.Box.Freq+1)
		return
	}
	if this.size == this.capacity {
		lastLruNode := this.tail.Pre.Tail.Pre
		delete(this.hashMap, lastLruNode.Key)
		lastLruNode.Key = key
		lastLruNode.Val = value
		this.removeNodeFromList(lastLruNode)
		this.insertNodeIntoBoxAfterLfuNode(lastLruNode, this.tail.Pre, 0)
		this.hashMap[key] = lastLruNode
	} else {
		newLruNode := &LruNode{key, value, nil, nil, nil}
		this.hashMap[key] = newLruNode
		this.insertNodeIntoBoxAfterLfuNode(newLruNode, this.tail.Pre, 0)
		this.size++
	}
}

func (this *LFUCache) insertNodeIntoBoxAfterLfuNode(lruNode *LruNode, lfuNode *LfuNode, freq int) {
	if lfuNode.Freq == freq { // 插入该lfuNode
		lruNode.Next = lfuNode.Head.Next
		lruNode.Pre = lfuNode.Head
		lruNode.Next.Pre = lruNode
		lfuNode.Head.Next = lruNode
		lruNode.Box = lfuNode
	} else { // 在该lfuNode之后插入新节点
		lruHead := &LruNode{-1, -1, nil, lruNode, nil}
		lruTail := &LruNode{-1, -1, lruNode, nil, nil}
		lruNode.Pre = lruHead
		lruNode.Next = lruTail
		newLfuNode := &LfuNode{freq, lfuNode, lfuNode.Next, lruHead, lruTail}
		lfuNode.Next.Pre = newLfuNode
		lfuNode.Next = newLfuNode
		lruHead.Box = newLfuNode
		lruTail.Box = newLfuNode
		lruNode.Box = newLfuNode
	}
}

/**
将lruNode从lru链表中移除
移除后，Box为空，则将Box也从lfu链表中移除
返回lfuNode
*/
func (this *LFUCache) removeNodeFromList(lruNode *LruNode) {
	this.removeNodeFromLruList(lruNode)
	if this.isLfuNodeEmpty(lruNode.Box) {
		this.removeNodeFromLfuList(lruNode.Box)
	}
}

func (this *LFUCache) removeNodeFromLruList(lruNode *LruNode) {
	lruNode.Pre.Next = lruNode.Next
	lruNode.Next.Pre = lruNode.Pre
}

func (this *LFUCache) removeNodeFromLfuList(lfuNode *LfuNode) {
	lfuNode.Next.Pre = lfuNode.Pre
	lfuNode.Pre.Next = lfuNode.Next
}

func (this *LFUCache) isLfuNodeEmpty(lfuNode *LfuNode) bool {
	return nil == lfuNode.Head.Next.Next
}
