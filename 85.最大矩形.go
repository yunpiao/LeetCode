/*
 * @lc app=leetcode.cn id=85 lang=golang
 * @lcpr version=30200
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	stack := []int{}
	maxArea := 0
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			w := i - stack[len(stack)-1] - 1
			maxArea = max(maxArea, h*w)
		}
		stack = append(stack, i)
	}
	return maxArea
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1"]]\n
// @lcpr case=end

*/

