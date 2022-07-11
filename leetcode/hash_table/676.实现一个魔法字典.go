package leetcode

/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start
type MagicDictionary struct {
	m  map[int]map[byte]bool
	m2 map[string]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{map[int]map[byte]bool{}, map[string]bool{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for k := range dictionary {
		this.m2[dictionary[k]] = true
		for kk := range dictionary[k] {
			if _, ok := this.m[kk]; !ok {
				this.m[kk] = map[byte]bool{}
			}
			this.m[kk][dictionary[k][kk]] = true
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	sc := []byte(searchWord)
	var temp byte
	for k := range searchWord {
		if bm, ok := this.m[k]; !ok {
			return false
		} else {
			sc = []byte(searchWord)
			temp = sc[k]
			for b := range bm {
				if b != temp {
					sc[k] = b
					if this.m2[string(sc)] {
						return true
					}
				}
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

/*
	m的kv关系为：m[长度索引][字符]{是否存在} ,m2记录存在的字符串

	build时，将字典的各个位的数据都存进去，search的时候，依次从m的对应长度索引中取不同的字节出来替换，然后在m2中找
*/
