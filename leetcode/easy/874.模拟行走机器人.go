package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	direction := 1 //当前方向 1北 y++  3南 y-- 4西 i-- 2东 i++
	ht := map[string]bool{}
	for k := range obstacles {
		ht[strconv.Itoa(obstacles[k][0])+","+strconv.Itoa(obstacles[k][1])] = true
	}
	x, y := 0, 0
	maxDist := 0
	for k := range commands {
		if commands[k] >= 1 && commands[k] <= 9 {
			temp := commands[k]
			for temp > 0 {
				checkObstacle874(&x, &y, direction, ht)
				temp--
			}
		} else {
			changeDirection874(commands[k], &direction)
		}
		if x*x+y*y > maxDist { //过程中最大
			maxDist = x*x + y*y
		}
	}
	return maxDist
}

//调整方向
func changeDirection874(value int, direction *int) {
	if value == -2 { //左转
		if *direction == 1 {
			*direction = 4
		} else {
			*direction--
		}
	} else if value == -1 { //右转
		if *direction == 4 {
			*direction = 1
		} else {
			*direction++
		}
	}
}

//检测障碍物
func checkObstacle874(x, y *int, direction int, ht map[string]bool) {
	if direction == 1 { //1北 y++  3南 y-- 4西 i-- 2东 i++
		if !ht[strconv.Itoa(*x)+","+strconv.Itoa(*y+1)] {
			*y++
		}
	} else if direction == 2 {
		if !ht[strconv.Itoa(*x+1)+","+strconv.Itoa(*y)] {
			*x++
		}
	} else if direction == 3 {
		if !ht[strconv.Itoa(*x)+","+strconv.Itoa(*y-1)] {
			*y--
		}
	} else {
		if !ht[strconv.Itoa(*x-1)+","+strconv.Itoa(*y)] {
			*x--
		}
	}
}

// @lc code=end
