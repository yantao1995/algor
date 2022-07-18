package leetcode

/*
 * @lc app=leetcode.cn id=749 lang=golang
 *
 * [749] 隔离病毒
 */

// @lc code=start
func containVirus(isInfected [][]int) int {
	m, n := len(isInfected), len(isInfected[0])
	type kv struct {
		x, y int
	}
	countMap := map[kv][]kv{}    //从kv点开始的影响个数
	oriMap := map[kv][]kv{}      //原始病毒的区域
	distinctMap := map[kv]bool{} //单次查找，路径控制
	var dfs func(c kv, x, y int) //c初始单元格，xy当前单元格
	find := false                //用于判断是否重复,存在1包裹0的情况
	dfs = func(c kv, x, y int) {
		if x >= 0 && x < m &&
			y >= 0 && y < n && !distinctMap[kv{x, y}] {
			if isInfected[x][y] == 0 {
				find = false
				for k := range countMap[c] {
					if countMap[c][k].x == x && countMap[c][k].y == y {
						find = true
						break
					}
				}
				if !find {
					countMap[c] = append(countMap[c], kv{x, y})
				}
			} else if isInfected[x][y] == 1 {
				distinctMap[kv{x, y}] = true
				oriMap[c] = append(oriMap[c], kv{x, y})
				dfs(c, x-1, y)
				dfs(c, x+1, y)
				dfs(c, x, y-1)
				dfs(c, x, y+1)
			}
		}
	}
	mOri := map[kv]bool{}
	getFirewall := func(ori, noVirus []kv) int {
		mOri = map[kv]bool{}
		for k := range ori {
			mOri[ori[k]] = true
		}
		x, y := 0, 0
		count := 0
		for k := range noVirus {
			x, y = noVirus[k].x, noVirus[k].y
			if x-1 >= 0 && isInfected[x-1][y] == 1 && mOri[kv{x - 1, y}] {
				count++
			}
			if x+1 < m && isInfected[x+1][y] == 1 && mOri[kv{x + 1, y}] {
				count++
			}
			if y-1 >= 0 && isInfected[x][y-1] == 1 && mOri[kv{x, y - 1}] {
				count++
			}
			if y+1 < n && isInfected[x][y+1] == 1 && mOri[kv{x, y + 1}] {
				count++
			}
		}

		for k := range ori {
			isInfected[ori[k].x][ori[k].y] = -1
		}
		return count
	}

	result := 0
	maxCount, maxCountKV := 0, kv{}
	for {
		distinctMap = map[kv]bool{}
		countMap = map[kv][]kv{}
		oriMap = map[kv][]kv{}
		maxCountKV = kv{}
		maxCount = 0
		for i := 0; i < m; i++ { //寻找感染区域
			for j := 0; j < n; j++ {
				if isInfected[i][j] == 1 && !distinctMap[kv{i, j}] {
					dfs(kv{i, j}, i, j)
				}
			}
		}
		for k := range countMap { //找最大感染区域
			if len(countMap[k]) > maxCount {
				maxCount = len(countMap[k])
				maxCountKV = k
			}
		}
		result += getFirewall(oriMap[maxCountKV], countMap[maxCountKV]) //计算防火墙个数，隔离感染区域
		for k := range countMap {                                       //感染其他区域
			if k != maxCountKV {
				for v := range countMap[k] {
					isInfected[countMap[k][v].x][countMap[k][v].y] = 1
				}
			}
		}
		if len(countMap) <= 1 { //上一轮找到的数量小于等于1就可以退出了
			return result
		}
	}
}

// @lc code=end

/*
	定义 kv 结构，用于在map中存储坐标值
	countMap 从kv点开始的会受到感染的所有坐标
	oriMap   从kv点开始的所有已经感染的坐标
	distinctMap 用于dfs路径控制，防止往回走
	dfs 用于找当前会受到感染的区域

	mOri 仅在 getFirewall 中，用于便捷找到当前未感染区域是因为受到当前的连续感染区域而感染的
		例如下面这种情况：
				0  0  0  0
				1  0  1  1
				0  0  0  0
			首先，程序会判断是两块区域，如果当前需要感染右边的区域，
			那么在判断坐标[1,1]位置时0的左边和右边的时候，会受到左边感染的区域的影响，
			而count++两次，所以在需要用mOri判断当前的感染判断是否是受到原始感染区域影响
			进行4个方向的判断是因为有这种感染区域包裹的情况，就需要大于1面的防火墙
			例子,下面这种情况就需要3面防火墙
				1 1
				0 1
				1 1

	getFirewall 用于找到当前区域设置的防火墙个数 并且对感染区域进行封锁，即对感染区域设置值为-1

	maxCount, maxCountKV 分别用于找到受影响最大的未感染区域 和 起始位置坐标



	步骤：
		1. 找到各个感染区域和被感染区域
		2. 找到最大的感染区域
		3. 根据感染区域计算防火墙的个数，然后隔离感染区域
		4. 感染下一轮区域
		5. 判断退出条件
*/
