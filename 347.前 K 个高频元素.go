/*
 * @lc app=leetcode.cn id=347 lang=golang
 * @lcpr version=30104
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"container/heap"
	"fmt"
)

// @lc code=start

type Item struct{ num, count int }
type MinHeap []Item

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].count < (*h)[j].count }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(Item)) }
func (h *MinHeap) Pop() interface{} {
	ret := (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	return ret
}
func topKFrequent(nums []int, k int) []int {
	// m := make(map[int]int, len(nums))
	// ret := make([]int, 0, len(nums))

	// for _, n := range nums {
	// 	m[n]++
	// }

	// // 使用自定义结构排序
	// type kv struct{ k, v int }
	// arr := make([]kv, 0, len(m))
	// for k, v := range m {
	// 	arr = append(arr, kv{k, v})
	// }
	// sort.Slice(arr, func(i, j int) bool { return arr[i].v > arr[j].v })

	// for i := 0; i < k; i++ {
	// 	ret = append(ret, arr[i].k)
	// }
	// return ret

	// 更优算法, 使用 heap 最小堆(顶出所有的最小元素, 剩下的就是最大的)

	d := make(map[int]int, len(nums))
	for _, n := range nums {
		d[n]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for i, v := range d {
		heap.Push(h, Item{i, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := []int{}
	for h.Len() != 0 {
		tmp := heap.Pop(h)
		ret = append(ret, tmp.(Item).num)
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
