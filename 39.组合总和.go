/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30100
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	// 回溯法
	if len(candidates) == 0 {
		return [][]int{}
	}
	ret := [][]int{}
	var dfs func(index int, path []int, sum int)
	dfs = func(index int, path []int, sum int) {
		if sum == target {
			ret = append(ret, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		// 允许出现逆序
		// for i := start; i < len(candidates); i++ { // index 为起点
		// 	dfs(i, append(path, candidates[i]), sum+candidates[i])
		// }

		// 不允许出现逆序 集合来求组合问题, 需要 start index
		// 这里 start 改成 0 就是所有的组合数
		for i := index; i < len(candidates); i++ { // index 为起点
			dfs(i, append(path, candidates[i]), sum+candidates[i])
		}
	}
	dfs(0, []int{}, 0)
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

