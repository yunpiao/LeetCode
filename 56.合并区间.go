/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=30104
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 1. 先排序, 按照左边为排序依据
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := [][]int{}
	if len(intervals) == 0 {
		return ret
	}
	// 2. for, 循环, 查看数据

	// 第一个区间, 取收元素的头部和尾部
	end := intervals[0][1]
	start := intervals[0][0]

	for _, v := range intervals {
		if v[0] > end {
			ret = append(ret, []int{start, end})
			start = v[0]
			end = v[1]
		}
		end = max(end, v[1])

	}
	ret = append(ret, []int{start, end})
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/

