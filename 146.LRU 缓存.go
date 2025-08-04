/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=30104
 *
 * [146] LRU 缓存
 */
package main

// @lc code=start
type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	// 双向链表 + map
	H *Node // 头链表
	T *Node // 尾部节点
	M map[int]*Node
	C int // 容量
}

// 关键点, 初始化并固定头尾指针
func Constructor(capacity int) LRUCache {
	h := &Node{}
	t := &Node{}
	h.Next = t
	t.Pre = h
	return LRUCache{
		H: h,
		T: t,
		M: make(map[int]*Node, capacity),
		C: capacity,
	}

}

// 获取的时候, 先删除后添加
func (this *LRUCache) Get(key int) int {
	if node, ok := this.M[key]; ok {
		// 已经是最后一位了, 直接返回 v
		if node.Next == nil {
			return node.Val
		}

		this.removeNode(node)
		this.addNode(node)
		return node.Val
	}
	return -1
}

// 添加的时候添加到尾部, 需要更改 T.Pre.Next  和 T.Pre
func (this *LRUCache) addNode(node *Node) {
	if node == nil {
		return
	}
	node.Pre = this.T.Pre
	node.Next = this.T
	this.T.Pre.Next = node
	this.T.Pre = node

	this.M[node.Key] = node
}

// 删除需要把 node.Pre.Next 和 node.Next.Pre 更新下, 因为有头尾节点, 所以没有尾部的情况
func (this *LRUCache) removeNode(node *Node) {
	if node == nil {
		return
	}
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	delete(this.M, node.Key)
}

// 如果存在, 先删除 后添加, 否则直接增加
func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.M[key]; ok {
		this.removeNode(node)
		node.Val = value
		this.addNode(node)
		return
	}

	// 容量检查, 发现超过容量, 头节点删除
	if len(this.M) == this.C {
		this.removeNode(this.H.Next)
	}
	this.addNode(&Node{Key: key, Val: value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
