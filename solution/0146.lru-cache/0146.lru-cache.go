package _146_lru_cache

// 146. LRU 缓存：中等
// tags：设计、哈希表、链表

type LRUCache struct {
	capacity int               // 容量
	knMap    map[int]*Node     // key -> Node(key, val)
	cache    *DoubleLinkedList // Node(k1, v1) <-> Node(k2, v2)...
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		knMap:    map[int]*Node{},
		cache:    initDoubleLinkedList(),
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.knMap[key]
	if !exist {
		// 未找到，返回 -1
		return -1
	}
	// 找到，将其设为最近使用
	this.makeRecently(key)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	_, exist := this.knMap[key]
	// 如果存在
	if exist {
		// 删除旧元素
		this.deleteKey(key)
		// 添加新数据为最近使用
		this.addRecently(key, value)
	} else {
		// 不存在
		// 容量满
		if this.cache.size == this.capacity {
			// 删除最近最少使用的元素
			this.deleteLRU()
		}
		// 添加元素为最近使用
		this.addRecently(key, value)
	}
}

// 将某个 key 调整为最近使用的元素
func (this *LRUCache) makeRecently(key int) {
	// 找到节点
	node := this.knMap[key]
	// 删除节点
	this.cache.remove(node)
	// 添加到链表尾部
	this.cache.addLast(node)
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key int, val int) {
	// 初始化节点
	node := initNode(key, val)
	// 添加到链表尾
	this.cache.addLast(node)
	// 别忘了关联map映射
	this.knMap[key] = node
}

// 删除某一个 key 及对应元素
func (this *LRUCache) deleteKey(key int) {
	node := this.knMap[key]
	// 删除节点
	this.cache.remove(node)
	// 删除映射
	delete(this.knMap, key)
}

// 删除最近最少使用的元素
func (this *LRUCache) deleteLRU() {
	// 链表的第一个就是最近最少使用的元素
	deletedNode := this.cache.removeFirst()
	// 删除映射
	delete(this.knMap, deletedNode.key)
}


// 链表节点
type Node struct {
	key, val   int   // 键、值
	prev, next *Node // 前、后
}

func initNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

// 双向链表
type DoubleLinkedList struct {
	head, tail *Node // 虚拟头节点、虚拟尾节点
	size       int   // 链表元素数
}

func initDoubleLinkedList() *DoubleLinkedList {
	dll := &DoubleLinkedList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
		size: 0,
	}
	dll.head.next = dll.tail
	dll.tail.prev = dll.head
	return dll
}

// 在链表尾部添加 x 节点，时间 O(1)
func (this *DoubleLinkedList) addLast(x *Node) {
	x.prev = this.tail.prev
	x.next = this.tail
	// 注意！上面两行和下面两行不能颠倒
	this.tail.prev.next = x
	this.tail.prev = x
	this.size++
}

// 删除链表中的 x 节点，时间 O(1)
func (this *DoubleLinkedList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	x.prev = nil
	x.next = nil
	this.size--
}

// 删除链表中的第一个节点，并返回该节点，时间 O(1)
func (this *DoubleLinkedList) removeFirst() *Node {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.remove(first)
	return first
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
