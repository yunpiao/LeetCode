/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=30005
 *
 * [139] 单词拆分
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// 动态规划,
	// dp 定义, dpi, 当前结尾是不是可以被组成
	// dp 初始化 0
	// dp 递归 dp[i] = dp[j] + nums[j:i] 属于 workDict, 存在最后一个词组可以拿来组成最后的结果
	n := len(s)
	if n == 0 {
		return true
	}
	dp := make([]bool, len(s)+1)
	dicts := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		dicts[v] = true
	}

	dp[0] = true
	// 这里也是类似的背包问题, 也是无限的子弹, 每个包的空间都计算之前的空间加上当前空间能不能拼成, 能的话就加上, 直接加 所以是无限子弹
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ { // 这里 j < i, 因为数据 s[j:i] 这个是不能取空的
			if dp[j] && dicts[s[j:i]] { // 之前是 true, 子数组在词典里面
				dp[i] = true // 这个是 true
				break
			}
		}
	}
	// 返回最后的数据
	return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n["leet", "code"]\n
// @lcpr case=end

// @lcpr case=start
// "applepenapple"\n["apple", "pen"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats", "dog", "sand", "and", "cat"]\n
// @lcpr case=end

*/

