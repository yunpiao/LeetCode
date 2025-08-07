/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=30201
 *
 * [6] Z 字形变换
 */
package main

// @lc code=start
func convert(s string, numRows int) string {
    buffer := make([][]byte, numRows)
    if numRows == 1 {
        return s
    }

    index := 0
    direct := 1

    for _, v := range s {
        // 如果是向下
        buffer[index] = append(buffer[index], byte(v))

        index += direct
        if index == numRows - 1 || index == 0 {
            direct = direct * -1
         }
    }

    ret := []byte("")
    for _,v := range buffer {
        ret = append(ret, v...)
    }
    return string(ret)
    
}
// @lc code=end



/*
// @lcpr case=start
// "PAYPALISHIRING"\n3\n
// @lcpr case=end

// @lcpr case=start
// "PAYPALISHIRING"\n4\n
// @lcpr case=end

// @lcpr case=start
// "AB"\n1\n
// @lcpr case=end

 */

