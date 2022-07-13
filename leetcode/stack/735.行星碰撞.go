package leetcode

/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	abs := 0
	stack := []int{}
lab:
	for k := range asteroids {
		if asteroids[k] > 0 {
			stack = append(stack, asteroids[k])
		} else {
			abs = 0 - asteroids[k]
			for len(stack) > 0 && stack[len(stack)-1] > 0 {
				if stack[len(stack)-1] >= abs {
					if stack[len(stack)-1] == abs {
						stack = stack[:len(stack)-1]
					}
					continue lab
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, asteroids[k])
		}
	}
	return stack
}

// @lc code=end

/*
	解1：
	stack 存储当前的行星，如果当前大于0，是向右的，可以直接存进去，
			如果小于0，那么是会向左碰撞，那么直接和栈顶的碰撞，
			直到这个向左碰撞的行星破碎，或者栈空时这个行星入栈。


	解2：
	直接模拟整个碰撞过程，找到 i-1 大于0   i小于0的位置开始碰撞
	然后判断当前位置，再继续碰撞
	func asteroidCollision(asteroids []int) []int {
	abs := 0
	for i := 1; i < len(asteroids); i++ {
		if i > 0 && asteroids[i] < 0 && asteroids[i-1] > 0 {
			abs = 0 - asteroids[i]
			if asteroids[i-1] > abs {
				asteroids = append(asteroids[:i], asteroids[i+1:]...)
				i--
			} else if asteroids[i-1] == abs {
				asteroids = append(asteroids[:i-1], asteroids[i+1:]...)
				i -= 2
			} else {
				asteroids = append(asteroids[:i-1], asteroids[i:]...)
				i -= 2
			}
		}
	}
	return asteroids
}
*/
