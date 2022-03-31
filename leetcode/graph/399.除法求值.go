package leetcode

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	rm := map[string]map[string]float64{} //元素/关联元素
	for index, eq := range equations {
		if _, ok := rm[eq[0]]; !ok {
			rm[eq[0]] = map[string]float64{}
		}
		rm[eq[0]][eq[1]] = values[index]
		rm[eq[0]][eq[0]] = 1
		if _, ok := rm[eq[1]]; !ok {
			rm[eq[1]] = map[string]float64{}
		}
		rm[eq[1]][eq[1]] = 1
		rm[eq[1]][eq[0]] = 1 / values[index]
	}
	//fmt.Println(rm)
	var graphFunc func(prev string, tails map[string]float64)
	graphFunc = func(prev string, tails map[string]float64) {
		for tail := range tails {
			for tail2, val2 := range rm[tail] {
				if prev != tail2 {
					// if prev == "x1" {
					// 	fmt.Println("prev", prev, "tail", tail, "val", rm[prev][tail], "tail2", tail2, "val2", val2, prev+"-"+tail2, rm[prev][tail2])
					// }
					if rm[prev][tail2] <= 0 {
						rm[prev][tail2] = rm[prev][tail] * val2
						//fmt.Println("_______", prev+"-"+tail2, rm[prev][tail2])
						graphFunc(prev, rm[tail2])
					}
				}
			}
		}
	}
	for prev, tails := range rm {
		graphFunc(prev, tails)
	}
	//
	// for prev, tails := range rm {
	// 	fmt.Println(prev, tails)
	// }
	result := make([]float64, len(queries))
	for i, query := range queries {
		if rm[query[0]][query[1]] > 0 {
			result[i] = rm[query[0]][query[1]]
		} else {
			result[i] = -1.0
		}
	}
	return result
}

// @lc code=end

/*
	rm  : relation map 表示 字符串变量 之间的关系,可以理解为邻接矩阵
	 rm[a][b:123] 表示  a/b值为 123
	 由 a/b 可以计算出  a/b  a/a  b/b  b/a 四个关系

	 若  rm[a][b:123,c:456]
	 	 rm[b][e:123,f:123]
		 rm[c].......
	 那么由  b的关系，即可以得到  a 与 e f 的关系
	 使用 graphFunc(prev,tails) 初始遍历时 prev = a   tails = [b,c]
	再次向下递归时，prev =a  tails 则为 rm[b],rm[c] 的子map,即获取了 a与 e f ...的关系
*/
