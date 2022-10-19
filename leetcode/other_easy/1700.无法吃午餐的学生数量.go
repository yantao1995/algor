package leetcode

/*
 * @lc app=leetcode.cn id=1700 lang=golang
 *
 * [1700] 无法吃午餐的学生数量
 */

// @lc code=start
func countStudents(students []int, sandwiches []int) int {
	count := 0
	for len(students) > 0 && count != len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			count = 0
		} else {
			count++
			students = append(students, students[0])
			students = students[1:]
		}
	}
	return len(students)
}

// @lc code=end

/*
	直接根据题意模拟整个过程即可
*/
