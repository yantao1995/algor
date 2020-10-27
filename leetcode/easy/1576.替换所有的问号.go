package easy

/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] 替换所有的问号
 */

// @lc code=start
func modifyString(s string) string {
	bts := []byte(s)
	for k := range bts {
		if bts[k] == '?' {
			bts[k] = 'a'
			head, tail := false, false
			if k == 0 {
				head = true
			}
			if k == len(bts)-1 {
				tail = true
			}
			if head && tail {
				continue
			} else if head && !tail {
				for bts[k] == bts[k+1] {
					bts[k]++
				}
			} else if !head && tail {
				for bts[k] == bts[k-1] {
					bts[k]++
				}
			} else {
				for bts[k] == bts[k-1] || bts[k] == bts[k+1] {
					bts[k]++
				}
			}
		}
	}
	return string(bts)
}

// @lc code=end
