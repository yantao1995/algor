package leetcode

/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */

// @lc code=start
// type Solution struct {
// 	regions     [][3]int   //[start,end)左闭右开区间,起始权值
// 	rd          *rand.Rand //随机数
// 	totalWeight int        //总权重
// }

// func Constructor(n int, blacklist []int) Solution {
// 	sort.Ints(blacklist)
// 	s := Solution{
// 		regions:     [][3]int{},
// 		rd:          rand.New(rand.NewSource(time.Now().UnixNano())),
// 		totalWeight: 0,
// 	}
// 	if len(blacklist) > 0 {
// 		if blacklist[0] > 0 {
// 			s.regions = append(s.regions, [3]int{0, blacklist[0], s.totalWeight})
// 			s.totalWeight += blacklist[0]
// 		}
// 	} else {
// 		s.regions = append(s.regions, [3]int{0, n, n})
// 		s.totalWeight += s.regions[len(s.regions)-1][2]
// 	}

// 	for i := 1; i < len(blacklist); i++ {
// 		if blacklist[i] > blacklist[i-1]+1 {
// 			s.regions = append(s.regions, [3]int{blacklist[i-1] + 1, blacklist[i], s.totalWeight})
// 			s.totalWeight += blacklist[i] - blacklist[i-1] - 1
// 		}
// 	}

// 	if len(blacklist) > 0 && blacklist[len(blacklist)-1] < n-1 {
// 		s.regions = append(s.regions, [3]int{blacklist[len(blacklist)-1] + 1, n, s.totalWeight})
// 		s.totalWeight += n - 1 - blacklist[len(blacklist)-1]
// 	}
// 	return s
// }

// func (this *Solution) Pick() int {
// 	if this.totalWeight > 0 {
// 		regionIndex := this.rd.Intn(this.totalWeight)
// 		for k := range this.regions {
// 			if k == len(this.regions)-1 ||
// 				(regionIndex >= this.regions[k][2] && regionIndex < this.regions[k+1][2]) {
// 				return this.regions[k][0] + this.rd.Intn(this.regions[k][1]-this.regions[k][0])
// 			}
// 		}
// 	}
// 	return 0
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end

/*
	思路：用黑名单数字将所有数字划分为各个区间，然后根据各个区间的可选择的数字个数，进行加权
		 然后对总权重进行随机，随机数落到对应的区间。在区间内的各个数字机会均等，然后再随机
		 到具体的数字

	regions 记录开闭区间和权重起始位置
*/
