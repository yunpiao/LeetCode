/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30003
 *
 * [3] 无重复字符的最长子串
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// todo, 滑动窗口题型
	// 核心数据结构：map记录字符最后出现的索引
	// 算法要点：
	// 1. 窗口维护：遇到重复字符时，左边界取 max(当前左边界, 重复字符上次位置+1)
	// 2. 边界处理：无论是否重复，每次右移都要更新字符索引和最大值
	// 注意：字符存在但索引小于左边界时，视为新字符，仍需更新索引位置
    n := len(s)
	if n == 0 {
		return 0
	}
	// 初始化
	l, r := 0, 0
	winMap := make(map[byte]int)
	ret := 0 
	// 滑动窗口
	for r < n {
		/// 如果之前不存在的字符, 当前索引加入到 map 中
		if _, ok := winMap[s[r]]; !ok {
			winMap[s[r]] = r
			// 更新最大值, 当前最大值和当前窗口最大值的比较
		}else {
			if winMap[s[r]] >= l {
				l = winMap[s[r]] + 1
			}
			// 更新重复字符的索引
			winMap[s[r]] = r
		}
		// 每次窗口更新, 都需要计算
		ret = max(ret, r-l+1)
		// 右边界右移
		r++
	}
	return ret

}
// @lc code=end



/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

 */

