/*
 * @lc app=leetcode.cn id=40 lang=golang
 * @lcpr version=30100
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates) // 排序, 用于去重
	var dfs func(start int, path []int, sum int)
	dfs = func(start int, path []int, sum int) {
		if sum == target {
			ret = append(ret, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		// start 如果改成 0, 就是所有的排列, 否则就是组合
		for i := start; i < len(candidates); i++ {
			// 如果发现, 横向拓展的时候, 之前已经取到过, 就不再重复取(结果去重, 因为元素可能是重复的)
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// 纵向拓展, i+1 表示, 每次只取一次
131		}
	}
	dfs(0, []int{}, 0)
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [10,1,2,7,6,1,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2,5,2,1,2]\n5\n
// @lcpr case=end

*/

