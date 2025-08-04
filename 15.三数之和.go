/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30005
 *
 * [15] 三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}

	for i1, v1 := range nums[:len(nums)-2] {
		// 如果没有负数, 不可能满足条件
		if i1 == 0 && nums[0] > 0 {
			return [][]int{}
		}
		// 第一个数字去重
		if i1 > 0 && nums[i1] == nums[i1-1] {
			continue
		}

		twoSumResults, ok := twoSum(nums[i1+1:], -v1)
		if ok {
			for _, twoSumResult := range twoSumResults {
				ret = append(ret, []int{v1, twoSumResult[0], twoSumResult[1]})
			}
		}
	}
	return ret

}

// 双指针法
func twoSum(nums []int, target int) ([][]int, bool) {
	n := len(nums)
	if n < 2 {
		return [][]int{}, false
	}
	ret := [][]int{}
	l, r := 0, n-1
	for l < r {
		if nums[l]+nums[r] == target {
			ret = append(ret, []int{nums[l], nums[r]})
			// 第 2 3 位去重
			for l < r && nums[l] == nums[l+1] {
				l++
			}
			for l < r && nums[r] == nums[r-1] {
				r--
			}
			l++
			r--
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			r--
		}
	}
	return ret, true
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/

