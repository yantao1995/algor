package leetcode

/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

// @lc code=start
func defangIPaddr(address string) string {
	count := 0
	for k := range address {
		if address[k] == '.' {
			count++
		}
	}
	if count > 0 {
		newAddress := make([]byte, 0, len(address)+count*2)
		for k := range address {
			if address[k] == '.' {
				newAddress = append(newAddress, '[')
			}
			newAddress = append(newAddress, address[k])
			if address[k] == '.' {
				newAddress = append(newAddress, ']')
			}
		}
		address = string(newAddress)
	}
	return address
}

// @lc code=end

// func defangIPaddr(address string) string {
// 	return strings.Replace(address, ".", "[.]", -1)
// }
