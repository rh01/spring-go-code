type ListNode struct {
	prev       *ListNode
	next       *ListNode
	key, value int
}

type LRUCache struct {
	head     *ListNode
	tail     *ListNode
	m        map[int]*ListNode
	capacity int
}

// 这个就是lru的基本操作，MoveToFront
func (lru *LRUCache) setHeader(node *ListNode) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

// remove操作就是移除掉最后的节点
func (node *ListNode) remove() {
	node.prev = node.next.prev
	node.prev.next = node.next
}

// 初始化LRU
func Constructure(capacity int) LRUCache {
	lru := LRUCache{
		head:     new(ListNode),
		tail:     new(ListNode),
		m:        make(map[int]*ListNode),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (lru LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; ok {
		val := node.value
		node.remove()
		lru.setHeader(node)
		return val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// 先检查是否存在
	if node, ok := lru.m[key]; ok {
		node.remove()
		node.value = value
		lru.setHeader(node)
	} else {
		if len(lru.m) == lru.capacity {
			node := lru.tail.prev
			key2 := node.key

			node.remove()
			delete(lru.m, key2)
		}
		node = &ListNode{key: key, value: value}
		lru.setHeader(node)
		lru.m[key] = node
	}

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
