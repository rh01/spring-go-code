package lcci

// LRUCache 常考可以看做是双向链表+map
type LRUCache struct {
	head *ListNode
	tail *ListNode

	cap int
	m   map[int]*ListNode
}

// ListNode 链表的节点
type ListNode struct {
	key, value int
	pre        *ListNode
	next       *ListNode
}

// 这里设置的是头结点的下一个节点，有时也被称为头节点
func (l *LRUCache) setHeader(node *ListNode) {
	l.head.next.pre = node
	node.next = l.head.next

	node.pre = l.head
	l.head.next = node
}

func (node *ListNode) remove() {
	node.pre = node.next.pre
	node.pre.next = node.next
}

// Constructor 构造一个新的缓存队列
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		head: new(ListNode),
		tail: new(ListNode),

		cap: capacity,
		m:   make(map[int]*ListNode, capacity),
	}
	lru.tail.pre = lru.head
	lru.head.next = lru.tail

	return lru
}

// Get 从LRU中获取一个key对应的值
func (l *LRUCache) Get(key int) int {
	// 如果获取到了
	if node, ok := l.m[key]; ok {
		value := node.value
		node.remove()
		l.setHeader(node)
		return value
	}
	return -1
}

// Put 将
func (l *LRUCache) Put(key int, value int) {
	// 首先查看当前的缓冲是否满了

	if node, ok := l.m[key]; ok {
		node.remove()
		delete(l.m, key)
	} else {
		if len(l.m) == l.cap {
			toRemove := l.tail.pre
			toRemove.remove()
			delete(l.m, toRemove.key)
		}
	}

	newNode := &ListNode{key: key, value: value}
	l.setHeader(newNode)
	l.m[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
