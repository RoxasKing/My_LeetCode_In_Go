package My_LeetCode_In_Go

/*
  运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
  获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
  写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

  进阶: 你是否可以在 O(1) 时间复杂度内完成这两种操作？
*/
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func (this *LRUCache) insert(node *Node) {
	tail := this.Tail
	node.Prev = tail.Prev
	tail.Prev.Next = node
	node.Next = tail
	tail.Prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

type LRUCache struct {
	limit int
	hash  map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := new(Node), new(Node)
	head.Next, tail.Prev = tail, head
	return LRUCache{
		limit: capacity,
		hash:  make(map[int]*Node, capacity),
		Head:  head,
		Tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Val = value
		return
	}
	if len(this.hash) >= this.limit {
		h := this.Head.Next
		this.remove(h)
		delete(this.hash, h.Key)
	}
	node := &Node{Key: key, Val: value}
	this.hash[key] = node
	this.insert(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */