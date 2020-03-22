package lcci

import "container/list"

// LRUCache 实现了最近最少使用缓存
type LRUCache struct {
	cap int                   // 缓存的容量
	l   *list.List            // 缓存队列，这里使用了list内置的List对象
	m   map[int]*list.Element // 利用哈希可以很快的找到列表中是否存在
}

// Pair 表示缓存中存放的东西
type Pair struct {
	key   int
	value int
}

// Constructor 构造一个最近最少使用缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get 获取指定的key对应的value
func (lru *LRUCache) Get(key int) int {
	// 检查是否存在对应的key
	if node, ok := lru.m[key]; ok {
		value := node.Value.(*list.Element).Value.(Pair).value
		// move node to front
		lru.l.MoveToFront(node)
		return value
	}

	return -1
}

// Put 讲一个key添加到缓存中，有两种情况，分别为
// 存在和不存子啊
func (lru *LRUCache) Put(key int, value int) {
	// 如果存在，更新即可
	if node, ok := lru.m[key]; ok {
		// 移动到最前面
		lru.l.MoveToFront(node)
		// 更新
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		// 如果当前的缓冲队列已满
		if lru.l.Len() == lru.cap {
			idx := lru.l.Back().Value.(*list.Element).Value.(Pair).key
			delete(lru.m, idx)
			lru.l.Remove(lru.l.Back())
		}

		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}

		ptr := lru.l.PushFront(node)
		lru.m[key] = ptr
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
