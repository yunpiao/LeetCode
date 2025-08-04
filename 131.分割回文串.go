/*
 * @lc app=leetcode.cn id=131 lang=golang
 * @lcpr version=30104
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	// 难在抽象概念, 这里用分割等于组合, 每次分割, 就是组合为 0:1 0:2 0:3 1:2 1:3 2:3 这种
	// 所以, 这里需要用 start 来表示, 每次分割的起点, 然后从 start 开始, 到 i 结束, 判断是否是回文
	// 如果是回文, 则将 s[start:i+1] 加入到 path 中, 然后继续分割 s[i+1:]
	// 如果不是回文, 则继续分割 s[start:i+1]
	// 直到 start >= len(s), 则将 path 加入到 ret 中
	// 最后返回 ret

	huiwen := func(s string) bool {
		if len(s) == 0 {
			return true
		}
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	if len(s) == 0 {
		return [][]string{}
	}
	var ret = [][]string{}

	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		if start >= len(s) {
			ret = append(ret, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if huiwen(s[start : i+1]) {
				dfs(i+1, append(path, s[start:i+1]))
			}
		}
	}
	dfs(0, []string{})
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/

