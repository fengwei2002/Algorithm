/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package leetcode

import "math"

/* --- 2022-01-07-22-57 --- */

// https://github.com/fengwei2002/Algorithm
// solution link:
// https://leetcode.cn/problems/reverse-integer/solution/lc7-fengwei2002-qiao-miao-chu-li-chao-gu-q0we/

// @lc code=start
func reverse(x int) int {
    ans := 0
    for x != 0 {
        // ans * 10 + x % 10 > INT_MAX
        if ans > 0 && ans > (math.MaxInt32 - x % 10) / 10 {
            return 0 
        }
        // ans * 10 + x % 10 < INT_MIN 
        if ans < 0 && ans < (math.MinInt32 - x % 10) / 10 {
            return 0
        }
        ans = ans * 10 + x % 10 
        x /= 10 
    }
    return int(ans)
}
// ans * 10 + x % 10 > INT_MAX == ans > (INT_MAX - x % 10) / 10

// @lc code=end
