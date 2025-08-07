/*
 * @lc app=leetcode.cn id=347 lang=golang
 * @lcpr version=30104
 *
 * [347] 前 K 个高频元素
 */
package main

import "container/heap"

// @lc code=start
type KV struct {
	k int
	v int
}

type MyHeap []KV
func (h MyHeap) Less(i, j int) bool {
	return h[i].v<h[j].v
}
func (h MyHeap) Len() int {
	return len(h)
}
func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MyHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *MyHeap) Push(v any) {
	*h = append(*h, v.(KV))
}

func topKFrequent(nums []int, k int) []int {
	h := &MyHeap{}
	heap.Init(h)
	set := make(map[int]int, len(nums))

	for _, v := range nums{
		set[v]++
	}
	for key,value := range set {
		if h.Len()<k{
			heap.Push(h, KV{key,value})
		} else if value>(*h)[0].v {// 大于栈顶的最小元素, 把栈顶的小元素删除, 当前元素加进去
			heap.Pop(h)
			heap.Push(h, KV{key, value})
		}
	}

	ret := make([]int, k)
	for i := range ret {
		ret[i] = heap.Pop(h).(KV).k
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
