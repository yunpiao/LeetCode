/*
 * @lc app=leetcode.cn id=216 lang=golang
 * @lcpr version=30104
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	// 1-9 中选k个数, 使得和为n
	// 回溯
	ret := [][]int{}
	var dfs func(start int, sum int, path []int)
	dfs = func(start int, sum int, path []int) {
		// 剪枝
		if len(path)+(9-start+1) < k {
			return
		}

		if len(path) == k {
			if sum == n {
				ret = append(ret, append([]int{}, path...))
			}
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)    // 添加当前数字
			dfs(i+1, sum+i, path)     // 递归
			path = path[:len(path)-1] // 回退
		}
	}
	dfs(1, 0, []int{})
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n9\n
// @lcpr case=end

// @lcpr case=start
// 4\n1\n
// @lcpr case=end

*/

