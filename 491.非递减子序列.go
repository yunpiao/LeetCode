/*
 * @lc app=leetcode.cn id=491 lang=golang
 * @lcpr version=30104
 *
 * [491] 非递减子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) == 0 {
		return ret
	}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 准入条件 至少有两个元素
		if len(path) >= 2 {
			ret = append(ret, append([]int{}, path...))
		}
		if start >= len(nums) {
			return
		}
		// 同一层级去重, 如果使用过, 就跳过
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			used[nums[i]] = true

			dfs(i+1, append(path, nums[i]))
		}

	}
	dfs(0, []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10,1,1,1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,4,3,2,1]\n
// @lcpr case=end

*/

