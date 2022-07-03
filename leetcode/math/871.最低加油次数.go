package leetcode

/*
 * @lc app=leetcode.cn id=871 lang=golang
 *
 * [871] 最低加油次数
 */

// @lc code=start
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	m := map[int]int{} //站 - 油
	getMaxM := func() int {
		max := -1
		if len(m) == 0 {
			return max
		}
		for k := range m {
			if m[k] > m[max] {
				max = k
			}
		}
		temp := m[max]
		delete(m, max)
		return temp
	}
	count := 0
	for i := 0; i < len(stations); i++ {
		for startFuel < stations[i][0] {
			if get := getMaxM(); get == -1 {
				return -1
			} else {
				startFuel += get
				count++
			}
		}
		if startFuel >= target {
			return count
		}
		m[i] = stations[i][1]
	}
	for startFuel < target {
		if get := getMaxM(); get == -1 {
			return -1
		} else {
			startFuel += get
			count++
		}
	}

	return count
}

// @lc code=end

/*
	存储当前油量能到达的所有站点，那么能到达的所有站点都可以把油量加到车里。
	因为站点是按距离升序排列，并且站点的距离小于目的地。
	所以当前油量不足以到下个站点或者目的地的时候，遍历已经走过的站点，
	因为要最少的次数，所以在未加的油里，挑一个最大油量的加
		如果没有油可加却仍没到下一个位置，那么不可达
		如果能到下一个位置，那么直接到下一个位置

	因为遍历的是 station ，所以遍历完成之后对目的地仍然要进行一次遍历

*/
