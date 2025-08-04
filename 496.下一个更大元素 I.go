/*
 * @lc app=leetcode.cn id=496 lang=golang
 * @lcpr version=30104
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 每日温度的多次调用版本

	// 用 map 存储 nums1 中元素值到索引的映射，方便快速查找
	dict := make(map[int]int, len(nums1))

	// 初始化结果, 默认是 -1
	ret := make([]int, len(nums1))
	for i := range nums1 {
		dict[nums1[i]] = i
		ret[i] = -1
	}

	// 单调栈
	stack := make([]int, 0, len(nums2))

	for i, v := range nums2 {
		for len(stack) > 0 && v > nums2[stack[len(stack)-1]] {
			nums2Index := stack[len(stack)-1]
			if nums1Index, ok := dict[nums2[nums2Index]]; ok {
				ret[nums1Index] = v
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,2]\n[1,3,4,2].\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n[1,2,3,4].\n
// @lcpr case=end

*/

