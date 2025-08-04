/*
 * @lc app=leetcode.cn id=312 lang=golang
 * @lcpr version=30104
 *
 * [312] 戳气球
 */
package main

// @lc code=start
func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 3; length <= n; length++ {
		for start := 0; start <= n-length; start++ {
			left := start
			right := left + length - 1
			for k := left + 1; k < right; k++ {
				dp[left][right] = max(dp[left][right], dp[left][k]+dp[k][right]+nums[left]*nums[right]*nums[k])
			}
		}
	}

	return dp[0][n-1]

	// 难一些的动态规划
	// 1. dp[i][j] 定义, 设置长度区间, 从 i 到 j 开区间之间的最大值
	// 2. 递推 dp[i][j] = max k->i~j ( dp[i][k]+dp[k][j]+n[i]*n[j]*n[k] )
	// 3. 确定遍历顺序. 长度相关的, 想从小的长度计算, 依次网上增加
	// 4. 打印输出

}

// @lc code=end

/*
// @lcpr case=start
// [3,1,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,5]\n
// @lcpr case=end

*/
