package leetcode

/*
 * @lc app=leetcode.cn id=2325 lang=golang
 *
 * [2325] 解密消息
 */

// @lc code=start
func decodeMessage(key string, message string) string {
	keyMap := map[byte]byte{}
	last := byte('a')
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && keyMap[key[i]] == 0 {
			keyMap[key[i]] = last
			last++
		}
	}
	result := []byte(message)
	for i := 0; i < len(result); i++ {
		if result[i] != ' ' {
			result[i] = keyMap[result[i]]
		}
	}
	return string(result)
}

// @lc code=end

/*
	使用keyMap建立映射关系，然后逐个替换即可
*/
