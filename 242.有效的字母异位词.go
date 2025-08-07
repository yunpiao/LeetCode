/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=30104
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}
	
	set := [26]int{}
	// 累加字符计数
	for _, v := range s {
		set[v-'a']++
	}

	// 逐步减少字符计数
	for _, v := range t {
		set[v-'a']--
		// 如果发现异常提前退出
		if set[v-'a'] < 0{
			return false
		}
	}
	
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/

