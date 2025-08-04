/*
 * @lc app=leetcode.cn id=707 lang=golang
 * @lcpr version=30104
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cursor := this.Next
	for i := 0; i < index && cursor != nil; i++ {
		cursor = cursor.Next
	}
	if cursor == nil {
		return -1
	}
	return cursor.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Next = &MyLinkedList{Next: this.Next, Val: val}
}

func (this *MyLinkedList) AddAtTail(val int) {
	cursor := this
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
	}

	cursor := this
	for i := 0; i < index && cursor != nil; i++ {
		cursor = cursor.Next
	}
	if cursor == nil {
		return
	}

	cursor.Next = &MyLinkedList{Val: val, Next: cursor.Next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	cursor := this
	for cursor.Next != nil && index != 0 {
		cursor = cursor.Next
		index--
	}
	if cursor.Next == nil {
		return
	}
	cursor.Next = cursor.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end



