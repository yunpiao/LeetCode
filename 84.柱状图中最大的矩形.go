/*
 * @lc app=leetcode.cn id=84 lang=golang
 * @lcpr version=30104
 *
 * [84] 柱状图中最大的矩形
 */
//  42. 接雨水 (opens new window)是找每个柱子左右两边第一个大于该柱子高度的柱子，而本题是找每个柱子左右两边第一个小于该柱子的柱子。
//
// 单调递增栈, 栈里面是单调递增, 当遇到小于栈顶的元素, 计算以当前元素为右边界, 栈顶元素为中间高度, 次栈顶为左边界的矩形,
// 也就是 栈顶是高, 左边边界是第一个小于这个高的边界, (宽是 右边界-左边界-1)

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 单调最大栈
	maxArea := 0
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack := make([]int, 0, len(heights))
	for i, v := range heights {
		for len(stack) > 0 && v < heights[stack[len(stack)-1]] {
			// 先取出当前值, 计算以当前值为高度的矩形的面积
			midIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 宽的取值有两种, 一种是中间节点, 使用栈顶元素到i之间作为宽, 高使用 midIndex(没有在栈里面和 i 和 midIndex 之间的都是大于 midIndex 的)
			// 一种是到头了, 直接使用 i 就可以
			w := 0
			// leftIndex 和 midIndex 之间  还有  midIndex 和 i  之间可能有值, 但是值一定是大于 hright[midIndex] 的, 所以 w = i - leftIndex - 1
			leftIndex := stack[len(stack)-1]
			w = i - leftIndex - 1

			h := heights[midIndex] // 当前矩形的高
			area := h * w
			maxArea = max(maxArea, area)
		}
		stack = append(stack, i)
	}
	return maxArea
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,5,6,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n
// @lcpr case=end

*/

