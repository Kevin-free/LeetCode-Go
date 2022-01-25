package _146_lru_cache

type LRUCache struct {
	capacity int
	knMap map[int]*Node
	cache *DobleLinkedList
}


func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		knMap: map[int]*Node{},
		cache: initDobleLinkedList(),
	}
	return lruCache
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.knMap[key]
	if !ok {
		return -1
	}
	// 找到，将其设为最近使用
	this.makeRecently(key)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	_, ok := this.knMap[key]
	// 如果存在
	if ok {
		this.deleteKey(key)
		this.addRecently(key, value)
	}else {
		// 不存在
		// 容量满
		if this.capacity == this.cache.size {
			// 删除最久未使用的
			this.deleteLRU()
		}
		this.addRecently(key, value)
	}
}

// 调整 key 为最近使用的元素
func (this *LRUCache) makeRecently(key int) {
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
func (this *LRUCache) deleteKey(key int){
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

// 链表结点
type Node struct {
	key, val int
	prev, next *Node
}

func initNode(key, val int) *Node {
	return &Node {
		key: key,
		val: val,
	}
}

// 双链表
type DobleLinkedList struct {
	head, tail *Node // 虚拟头，虚拟尾
	size int
}

func initDobleLinkedList() *DobleLinkedList {
	// noto:diff
	dll := &DobleLinkedList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
		size: 0,
	}
	dll.head.next = dll.tail
	dll.tail.prev = dll.head
	return dll
}

// 在链表尾部添加节点 x
func (this *DobleLinkedList) addLast(x *Node) {
	x.prev = this.tail.prev
	x.next = this.tail
	this.tail.prev = x
	this.tail.prev.next = x
	this.size++
}

// 在链表中删除节点 x
func (this *DobleLinkedList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	x.next = nil
	x.prev = nil
	this.size--
}

// 在链表中删除第一个节点
func (this *DobleLinkedList) removeFirst() *Node {
	if this.head.next == this.tail {
		return nil
	}
	firstNode := this.head.next
	this.remove(firstNode)
	return firstNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */