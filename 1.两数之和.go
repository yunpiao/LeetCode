/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30002
 *
 * [1] 两数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {
	// 需要一个 hash 表判断是不是当前需要的数据
	set := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, exist := set[n]; exist {
			return []int{i, j}
		} else {
			set[target-n] = i
		}
	}
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

