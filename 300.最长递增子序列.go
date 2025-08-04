/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30103
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// dp 的定义比较难, 定义 dp是以当前 i 结尾的递增序列
	// 最后再使用 max 得到最大的子序列长度
	// 初始化 dp 数组, 每个元素初始化为 1, 因为每个元素本身就是一个长度为 1 的递增序列

	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ret := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}

	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

