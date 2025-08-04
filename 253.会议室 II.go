// @lc code=start
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// minMeetingRooms 扫描线做法
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 遍历每个会议, 如果当前会议开始时间 >= 最早结束会议的结束时间, 则可以复用会议室
	// 也就是说, 堆里面就是目前正在进行的会议的结束时间, 如果发现冲突, 就把之前的去掉
	for _, iv := range intervals {
		if h.Len() > 0 && iv[0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, iv[1])
	}

	return h.Len()

}

func main() {
	if minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}) != 2 {
		fmt.Println("error 1")
	}
	if minMeetingRooms([][]int{{7, 10}, {2, 4}}) != 1 {
		fmt.Println("error 2")
	}
	if minMeetingRooms([][]int{{1, 5}, {8, 9}, {8, 9}}) != 2 {
		fmt.Println("error 3")
	}
	fmt.Println("ok")
}
