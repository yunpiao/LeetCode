/*
 * @lc app=leetcode.cn id=435 lang=golang
 * @lcpr version=30104
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	// 	记录无重叠区域有多少个
	// 1. 先排序. 左边排序
	// 2. 之后计算, 根据每一个区间的最小 end 计算

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 1                     // 默认已经有一个
	minInterval := intervals[0][1] // 初始化一个默认的
	for i, interval := range intervals {
		if i == 0 {
			continue
		}
		// 需要一个新的区间
		if interval[0] >= minInterval {
			count++ 
			minInterval = interval[1]
		} else {
			// 更新当前区间的右侧最小值
			minInterval = min(minInterval, interval[1])
		}
	}
	// 最小的值 - 有重合的区间数量
	return len(intervals) - count

}

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[2,3],[3,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [1,2], [1,2] ]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [2,3] ]\n
// @lcpr case=end

*/

