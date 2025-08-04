/*
 * @lc app=leetcode.cn id=96 lang=golang
 * @lcpr version=30005
 *
 * [96] 不同的二叉搜索树
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i:=2; i<=n; i++{
        dp[i] = 0
        for j:=0; j<i; j++{
            dp[i] += dp[j] * dp[i-j-1]
        }
    }
    return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

