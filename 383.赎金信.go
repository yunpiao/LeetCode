/*
 * @lc app=leetcode.cn id=383 lang=golang
 * @lcpr version=30104
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	set := [26]int{}

	for _, v := range magazine {
		set[v-'a']++
	}

	for _, v := range ransomNote {
		set[v-'a']--
		if set[v-'a'] < 0 {
			return false
		}
	}
	

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "a"\n"b"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"ab"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"aab"\n
// @lcpr case=end

*/

