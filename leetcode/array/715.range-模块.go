package leetcode

/*
 * @lc app=leetcode.cn id=715 lang=golang
 *
 * [715] Range 模块
 */

// @lc code=start
type RangeModule struct {
	rgs      [][2]int
	hasMixed func(a, b *[2]int) (int, int, bool) //左边界，右边级，是否有交集
}

// func Constructor() RangeModule {
// 	return RangeModule{
// 		rgs: [][2]int{},
// 		hasMixed: func(a, b *[2]int) (int, int, bool) {
// 			if a[0] < b[0] {
// 				a, b = b, a
// 			}
// 			if a[0] >= b[0] && a[0] <= b[1] {
// 				if a[1] > b[1] {
// 					return b[0], a[1], true
// 				}
// 				return b[0], b[1], true
// 			} else if a[1] >= b[0] && a[1] <= b[1] {
// 				if a[0] < b[0] {
// 					return a[0], b[1], true
// 				}
// 				return b[0], b[1], true
// 			}
// 			return 0, 0, false
// 		},
// 	}
// }

func (this *RangeModule) AddRange(left int, right int) {
	if len(this.rgs) == 0 ||
		left > this.rgs[len(this.rgs)-1][1] {
		this.rgs = append(this.rgs, [2]int{left, right})
	} else if left == this.rgs[len(this.rgs)-1][1] {
		this.rgs[len(this.rgs)-1][1] = right
	} else if right < this.rgs[0][0] {
		this.rgs = append([][2]int{{left, right}}, this.rgs...)
	} else if right == this.rgs[0][0] {
		this.rgs[0][0] = left
	} else {
		nn := [2]int{left, right}
		for i := 0; i < len(this.rgs); i++ {
			if nl, nr, ok := this.hasMixed(&nn, &this.rgs[i]); ok {
				this.rgs[i][0], this.rgs[i][1] = nl, nr
				for j := i + 1; j < len(this.rgs); {
					if nl, nr, ok := this.hasMixed(&this.rgs[i], &this.rgs[j]); ok {
						this.rgs[i][0], this.rgs[i][1] = nl, nr
						this.rgs = append(this.rgs[:j], this.rgs[j+1:]...)
					} else {
						break
					}
				}
				return
			} else if left > this.rgs[i][1] && i+1 < len(this.rgs) && right < this.rgs[i+1][0] {
				this.rgs = append(this.rgs[:i+1], append([][2]int{nn}, this.rgs[i+1:]...)...)
				return
			}
		}
	}
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	for i := 0; i < len(this.rgs); i++ {
		if left >= this.rgs[i][0] && right <= this.rgs[i][1] {
			return true
		}
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	for i := 0; i < len(this.rgs); i++ {
		if right < this.rgs[i][0] {
			return
		}
		if left < this.rgs[i][0] {
			if right <= this.rgs[i][0] {
				return
			} else if right > this.rgs[i][0] {
				if right < this.rgs[i][1] {
					this.rgs[i][0] = right
					return
				} else if right >= this.rgs[i][1] {
					this.rgs = append(this.rgs[:i], this.rgs[i+1:]...)
					i--
				}
			}
		} else if left == this.rgs[i][0] {
			if right < this.rgs[i][1] {
				this.rgs[i][0] = right
				return
			} else if right >= this.rgs[i][1] {
				this.rgs = append(this.rgs[:i], this.rgs[i+1:]...)
				i--
			}
		} else if left > this.rgs[i][0] {
			if left < this.rgs[i][1] {
				if right > this.rgs[i][1] {
					this.rgs[i][1] = left
				} else if right == this.rgs[i][1] {
					this.rgs[i][1] = left
					return
				} else if right < this.rgs[i][1] {
					nn := [2]int{right, this.rgs[i][1]}
					this.rgs[i][1] = left
					this.rgs = append(this.rgs[:i+1],
						append([][2]int{nn}, this.rgs[i+1:]...)...)
					return
				}
			}
		}
	}
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
// @lc code=end

/*
	总体思路：区间按左值排序，然后合并相交的区间
	比如 add  [1,2]  add   [2,3]   得到区间  [1,3]
	判断的条件比较多，容易忽略的一点是,不相交的区间的添加，比如在区间  [1,3] ,[6,7] 添加 [4,5] ,需要额外判断
*/
