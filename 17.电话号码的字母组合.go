/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=30005
 *
 * [17] 电话号码的字母组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitsMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ret := []string{}
	var backtrack func(path string, index int)
	backtrack = func(path string, index int) {
		// 结束条件, 如果长度满足, 则代表到达最后节点
		if index == len(digits) {
			ret = append(ret, path)
			return
		}
		// 得到当前 index 可能的字母
		for _, v := range digitsMap[digits[index]] {
			backtrack(path+string(v), index+1)
		}
	}
	backtrack("", 0)
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/

