/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=30104
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ret := [][]int{}

	sort.Ints(nums)

	var dfs func(used []bool, path []int)
	dfs = func(used []bool, path []int) {
		if len(path) == len(nums) {
			ret = append(ret, append([]int{}, path...))
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			// 横向去重, 要求必须大于 1, 必须以前使用过了
			// 也可以纵向去重复, used[i-1] == true, 就是纵向发现重复, 就不继续
			// 两个都能得到相应的结果
			// 对于排列问题，树层上去重和树枝上去重，都是可以的，但是树层上去重效率更高！
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			used[i] = true
			dfs(used, append(path, nums[i]))
			used[i] = false
		}

	}
	dfs(make([]bool, len(nums)), []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

