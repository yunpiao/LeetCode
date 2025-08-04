/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30104
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	ret := [][]int{}

	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 剪枝, 如果剩下的元素不够, 直接返回
		if len(path)+(n-start+1) < k {
			return
		}
		// 根据终止条件来控制树的深度
		if len(path) == k {
			ret = append(ret, append([]int{}, path...))
		}
		// 树的广度
		for i := start; i <= n; i++ { // startIndex来记录下一层递归，搜索的起始位置

			// 回溯, 先增加, 在撤销
			// path = append(path, i)
			// dfs(i+1, path)
			// path = path[:len(path)-1]

			// 也可以直接合成一句
			dfs(i+1, append(path, i))

		}

	}
	dfs(1, []int{})

	return ret

}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

