/*
 * @lc app=leetcode.cn id=93 lang=golang
 * @lcpr version=30104
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	// 剪纸 是不是一有效的 ip 地址 4 到 12 位
	// 判断每次是不是在 0 到 255 之间
	// 如果超过 一位不能开始是 0
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	var isValid = func(s string) bool {
		if len(s) == 0 {
			return false
		}
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		num := 0
		for _, v := range s {
			if v < '0' || v > '9' {
				return false
			}
			num = num*10 + int(v-'0')
		}
		if num > 255 || num < 0 {
			return false
		}
		return true
	}

	var ret = []string{}
	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		// 被分割成了四分, 而且到了最后一位, 就加到结果集合里
		if len(path) == 4 && start >= len(s) {
			ret = append(ret, strings.Join(path, "."))
			return
		} else if len(path) == 4 && start < len(s) {
			return
		}

		if start >= len(s) {
			return
		}

		for i := start; i < len(s); i++ {
			if isValid(s[start : i+1]) {
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
// "25525511135"\n
// @lcpr case=end

// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "101023"\n
// @lcpr case=end

*/

