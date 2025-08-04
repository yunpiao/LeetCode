/*
 * @lc app=leetcode.cn id=90 lang=golang
 * @lcpr version=30104
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	if len(nums) == 0 {
		return ret
	}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		ret = append(ret, append([]int{}, path...))
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			dfs(i+1, append(path, nums[i]))
		}

	}
	dfs(0, []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

