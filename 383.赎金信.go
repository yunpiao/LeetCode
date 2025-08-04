/*
 * @lc app=leetcode.cn id=383 lang=golang
 * @lcpr version=30104
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	letterCount := [26]int{}

	for _, m := range magazine {
		letterCount[m-'a']++
	}

	for _, r := range ransomNote {
		letterCount[r-'a']--
		if letterCount[r-'a'] < 0 {
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

