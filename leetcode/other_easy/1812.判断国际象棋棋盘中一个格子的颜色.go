package leetcode

/*
 * @lc app=leetcode.cn id=1812 lang=golang
 *
 * [1812] 判断国际象棋棋盘中一个格子的颜色
 */

// @lc code=start
func squareIsWhite(coordinates string) bool {
	return int(coordinates[0]-'a')%2 != int(coordinates[1]-'1')%2
}

// @lc code=end

/*
	直接观察棋盘，行与列对2取余不相等的为白色。
*/
