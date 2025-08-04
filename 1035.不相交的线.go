/*
 * @lc app=leetcode.cn id=1035 lang=golang
 * @lcpr version=30103
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {

			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 这里是不是可以取其他的
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]

}

// @lc code=end

/*
// @lcpr case=start
// [1,4,2]\n[1,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,5,1,2,5]\n[10,5,2,1,5,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,1,7,5]\n[1,9,2,5,1]\n
// @lcpr case=end

*/

