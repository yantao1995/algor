package leetcode

/*
 * @lc app=leetcode.cn id=1041 lang=golang
 *
 * [1041] 困于环中的机器人
 */

// @lc code=start
func isRobotBounded(instructions string) bool {
	//direct 1北 2东 3南 4西
	x, y, direct := 0, 0, 1
	for i := 0; i < 4; i++ {
		for _, v := range instructions {
			if v == 'G' {
				switch direct {
				case 1:
					y++
				case 2:
					x++
				case 3:
					y--
				case 4:
					x--
				}
			} else if v == 'L' {
				direct--
				if direct == 0 {
					direct = 4
				}
			} else {
				direct++
				if direct == 5 {
					direct = 1
				}
			}
		}
		if x == 0 && y == 0 && direct == 1 {
			return true
		}
	}
	return x == 0 && y == 0 && direct == 1
}

// @lc code=end

/*
	只要能在一个周期内回到回到初始状态，那么就说明是在环中
	一个周期的判断，即最大的周期为每次执行都在不同的象限中

	例：
		第一象限操作为: GRGRGRGL
		(0,0)->(0,1)->(1,1)->(1,0)->(0,0)
	依次向下最多4次就能回到初始状态
*/
