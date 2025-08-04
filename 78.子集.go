/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=30104
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// 还是回溯法, 与其他题不同, 这里不再取叶子节点, 而是取所有节点
	var ret = [][]int{}
	if len(nums) == 0 {
		return ret
	}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		ret = append(ret, append([]int{}, path...))
		if start >= len(nums) {
			return
		}
		// 因为下一层迭代使用的是 start, 所以传入的时候要往下走, 所以 dfs(i+1, append(path, nums[i]))
		for i := start; i < len(nums); i++ {
			dfs(i+1, append(path, nums[i]))
		}

	}
	dfs(0, []int{})
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

