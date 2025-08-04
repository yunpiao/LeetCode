/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=30104
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	// 全排列属于🌳状结构里面, 纵向里面不能重复用的场景,
	// 纵向去重使用 map 来记录已经用过的
	// 横向去重使用 for i := start + 1 来作为启动
	ret := [][]int{}
	if len(nums) == 0 {
		return ret
	}
	var dfs func(used []int, path []int)
	dfs = func(used []int, path []int) {
		if len(path) == len(nums) {
			ret = append(ret, append([]int{}, path...))
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			dfs(used, append(path, nums[i]))
			used[i] = 0
		}
	}
	dfs(make([]int, len(nums)), []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

