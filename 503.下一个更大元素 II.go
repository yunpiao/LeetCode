/*
 * @lc app=leetcode.cn id=503 lang=golang
 * @lcpr version=30104
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := make([]int, len(nums))
	for i := range ret {
		ret[i] = -1
	}

	stack := make([]int, 0, len(nums))
	for i := 0; i < len(nums)*2; i++ {
		index := i % len(nums)
		for len(stack) > 0 && nums[index] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 注意点 这里小在注意, 弹出来的元素是结果的索引
			ret[top] = nums[index]
		}
		stack = append(stack, index)
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,3]\n
// @lcpr case=end

*/

