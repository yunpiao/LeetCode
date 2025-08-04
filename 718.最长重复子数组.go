/*
 * @lc app=leetcode.cn id=718 lang=golang
 * @lcpr version=30103
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	// 定义一个二维dp, 记录每次的最大长度
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	ret := 0

	// 定义的时候多加一行和一列, 方便初始化
	// dp[0][x]
	// dp[x][0] 是没有意义的, 因为 0 代表没有

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			// 如果 nums[i-1]相等, 那更新 dp[i]
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ret = max(ret, dp[i][j])
		}
	}

	/* 优化空间复杂度, i 只依赖 i-1, 所以可以优化为 一维数组
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([]int, len(nums2)+1)
	ret := 0

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			// 如果 nums[i-1]相等, 那更新 dp[i]
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			}
			ret = max(ret, dp[j])
		}
	}
	return ret
	*/
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,2,1]\n[3,2,1,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n[0,0,0,0,0]\n
// @lcpr case=end

*/

