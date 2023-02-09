package leetcode

/*
 * @lc app=leetcode.cn id=1797 lang=golang
 *
 * [1797] 设计一个验证系统
 */

// @lc code=start
/* type AuthenticationManager struct {
	timeToLive int
	m          map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		m:          map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := this.m[tokenId]; ok && v+this.timeToLive > currentTime {
		this.m[tokenId] = currentTime
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	result := 0
	for _, v := range this.m {
		if v+this.timeToLive > currentTime {
			result++
		}
	}
	return result
}
*/
/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
// @lc code=end

/*
	timeToLive存过期时间
	map存当前时间
	每次比较判断即可
*/
