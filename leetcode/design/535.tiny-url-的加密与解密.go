package leetcode

/*
 * @lc app=leetcode.cn id=535 lang=golang
 *
 * [535] TinyURL 的加密与解密
 */

// @lc code=start
// type Codec struct {
// 	encoder base64.Encoding
// }

// func Constructor() Codec {
// 	return Codec{*base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")}
// }

// // Encodes a URL to a shortened URL.
// func (this *Codec) encode(longUrl string) string {
// 	return this.encoder.EncodeToString([]byte(longUrl))
// }

// // Decodes a shortened URL to its original URL.
// func (this *Codec) decode(shortUrl string) string {
// 	bts, _ := this.encoder.DecodeString(shortUrl)
// 	return string(bts)
// }

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
// @lc code=end

/*
	base64 来做加解密工具
*/
