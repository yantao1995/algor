package leetcode

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
// func grayCode(n int) []int {   //回溯 n>=10 超时
// 	getBit := func(i, n int) []int {
// 		rlt := make([]int, n)
// 		for idx := len(rlt) - 1; i > 0 || idx >= 0; idx, i = idx-1, i>>1 {
// 			rlt[idx] = i & 1
// 		}
// 		return rlt
// 	}
// 	cmp := func(a, b []int) bool {
// 		count := 0
// 		for i := 0; i < len(a) && count < 3; i++ {
// 			if a[i] != b[i] {
// 				count++
// 			}
// 		}
// 		return count == 1
// 	}
// 	bitMap := map[int][]int{}
// 	for i := 0; i < 1<<n; i++ {
// 		bitMap[i] = getBit(i, n)
// 	}
// 	relation := map[int][]int{}
// 	for i := 0; i < 1<<n; i++ {
// 		for j := 0; j < 1<<n; j++ {
// 			if i != j {
// 				if cmp(bitMap[i], bitMap[j]) {
// 					if _, ok := relation[i]; !ok {
// 						relation[i] = []int{}
// 					}
// 					relation[i] = append(relation[i], j)
// 				}
// 			}
// 		}
// 	}
// 	distinct := map[int]bool{0: true}
// 	result := []int{}
// 	find := false
// 	var backtrace func(ri int, temp []int)
// 	backtrace = func(ri int, temp []int) {
// 		if len(temp) == 1<<n || find {
// 			if !find && cmp(bitMap[temp[len(temp)-1]], bitMap[temp[0]]) {
// 				result = append(result, temp...)
// 				find = true
// 			}
// 			return
// 		}
// 		for _, next := range relation[ri] {
// 			if !distinct[next] {
// 				distinct[next] = true
// 				backtrace(next, append(temp, next))
// 				distinct[next] = false
// 			}
// 		}
// 	}
// 	backtrace(0, []int{0})
// 	return result
// }

func grayCode(n int) []int {
	stack := make([]int, 0, 1<<n)
	stack = append(stack, 0)
	for i := 1; i <= n; i++ { //在第几位加权1  ，右边第一个为1
		for tail := len(stack) - 1; tail >= 0; tail-- {
			stack = append(stack, stack[tail]+1<<(i-1))
		}
	}
	return stack
}

// @lc code=end

/*
	找规律 ，题目test case 示例2 的顺序的是在末尾补1，
	但是尾数多了之后，需要遍历每一个进制来补，比较麻烦。
	所以看case 示例1 的规律，直接在每位的首位补1。但是由于
	出现高位之后，比如在n次之后  0 变成了1000000
	所以最末尾应该是 那么首位是0的话，末尾就应该是  1000000
	所以在遍历结果集的时候，就应该从末尾处读取，就类似一个栈

	case 1
	[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
	- 00 和 01 有一位不同
	- 01 和 11 有一位不同
	- 11 和 10 有一位不同
	- 10 和 00 有一位不同
*/
