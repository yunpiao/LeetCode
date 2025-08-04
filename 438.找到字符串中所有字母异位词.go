/*
 * @lc app=leetcode.cn id=438 lang=golang
 * @lcpr version=30104
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	arr1 := [26]int{}
	arr2 := [26]int{}

	for i := 0; i < len(p); i++ {
		arr1[p[i]-'a']++
		arr2[s[i]-'a']++
	}

	ret := []int{}

	for i := len(p); i < len(s); i++ {
		if arr1 == arr2 {
			ret = append(ret, i-len(p))
		}
		arr2[s[i]-'a']++

		arr2[s[i-len(p)]-'a']--
	}
	if arr1 == arr2 {
		ret = append(ret, len(s)-len(p))
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/

