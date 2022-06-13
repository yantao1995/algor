package leetcode

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	clos := len(s) / 2
	dist := [][]byte{}
	for i := 0; i < numRows; i++ {
		dist = append(dist, []byte{})
		for j := 0; j <= clos; j++ {
			dist[i] = append(dist[i], []byte{' '}...)
		}
	}
	i, j, seq := 0, 0, true // i,j行列标记，true向下 向上false
	if numRows > 1 {        //因为先进行i++所以从-1开始
		i = -1
	}
	for k := 0; k < len(s); k++ {
		if seq {
			i++
		} else {
			i--
			j++
		}
		//	fmt.Println(k, i, j, (i == 0 && !seq), (i == numRows-1 && seq))
		dist[i][j] = s[k]
		if (i == 0 && !seq) || (i == numRows-1 && seq) {
			seq = !seq //调头
		}
	}
	// for i := 0; i < numRows; i++ { //输出格式
	// 	println()
	// 	for j := 0; j <= clos; j++ {
	// 		fmt.Printf("%s\t", string(dist[i][j]))
	// 	}
	// }
	rtn := []byte{}
	for i := 0; i < numRows; i++ {
		for j := 0; j <= clos; j++ {
			if dist[i][j] != ' ' {
				rtn = append(rtn, dist[i][j])
			}
		}
	}
	return string(rtn)
}

// @lc code=end
