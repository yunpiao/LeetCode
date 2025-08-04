/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=30104
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// 根据长度比较
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	// 通过字典统计每个字符的数据量
	letterCount := [26]int{}
	for _, c := range s {
		letterCount[c-'a']++
	}

	for _, c := range t {
		letterCount[c-'a']--
		// 提前返回
		if letterCount[c-'a'] < 0 {
			return false
		}
	}

	for _, count := range letterCount {
		if count != 0 {
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

