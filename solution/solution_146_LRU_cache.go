package solution

/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache

实例：

LRUCache cache = new LRUCache(2);

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4

*/

type DLinkedNode struct {
	key  int
	val  int
	prev *DLinkedNode
	next *DLinkedNode
}

type LRUCache struct {
	size     int
	capacity int
	head     *DLinkedNode
	tail     *DLinkedNode
	record   map[int]*DLinkedNode
}

func NewLRUCache(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
		record:   make(map[int]*DLinkedNode),
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	p, ok := this.record[key]
	if !ok {
		return -1
	}

	// 把p节点移至头部
	this.remove(p)
	this.addToHead(p)

	return p.val
}

func (this *LRUCache) Put(key int, value int) {
	if p, ok := this.record[key]; !ok {
		node := &DLinkedNode{key: key, val: value}

		// 满了,淘汰最后一个
		if this.size == this.capacity {
			last := this.tail.prev

			this.remove(last)

			delete(this.record, last.key)
			this.size--
		}

		// 加入头部
		this.addToHead(node)

		this.record[key] = node
		this.size++
	} else {
		// 移至头部
		this.remove(p)
		this.addToHead(p)

		p.val = value
	}

}

func (this *LRUCache) remove(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.next.prev = node
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) addToTail(node *DLinkedNode) {
	node.prev = this.tail.prev
	this.tail.prev.next = node
	node.next = this.tail
	this.tail.prev = node
}
