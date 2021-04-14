package main

import "fmt"

type LinkNode struct {
	key  int
	val  int
	prev *LinkNode // 头节点
	next *LinkNode // 尾结点
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	// 头尾都为空，做标记
	head, tail := &LinkNode{0, 0, nil, nil}, &LinkNode{0, 0, nil, nil}
	head.next, tail.prev = tail, head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	if v, exist := this.m[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) AddNode(node *LinkNode) {
	head := this.head
	node.next, head.next.prev = head.next, node
	node.prev, head.next = head, node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Put(key int, value int) {
	tail := this.tail
	if v, exist := this.m[key]; exist {
		v.val = value
		this.MoveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(this.m) == this.cap {
			delete(this.m, tail.prev.key)
			this.RemoveNode(tail.prev)
		}
		this.AddNode(v)
		this.m[key] = v
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
}
