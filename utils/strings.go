package utils

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

//GetStrSliceByLength string按长度分割字符串
func GetStrSliceByLength(src string, length int) []string {
	srcsli := strings.Split(src, "")
	rtnSli := make([]string, 0)
	temp := ""
	for i := 0; i < len(srcsli); i++ {
		temp += srcsli[i]
		if (i+1)%length == 0 && i != 0 {
			rtnSli = append(rtnSli, temp)
			temp = ""
		}
	}
	if temp != "" {
		rtnSli = append(rtnSli, temp)
	}
	return rtnSli
}

// NonceStrCreate 生成随机字符串 相邻不重复
func NonceStrCreate(len int) string {
	if len <= 0 {
		return ""
	}
	var bhint int
	var c, cc rune
	var buffer bytes.Buffer
	var length int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		if r.Intn(10) >= 5 {
			bhint = r.Intn(4) + r.Intn(23)
			c = rune(97 + bhint)
		} else {
			bhint = r.Intn(4) + r.Intn(23)
			c = rune(65 + bhint)
		}
		if length == 0 {
			buffer.WriteString(string(c))
			cc = c
			length = 1
		}
		if cc != c && length > 0 {
			buffer.WriteString(string(c))
			cc = c
			length++
		}
		if length == len {
			break
		}
	}
	return buffer.String()
}

//str1 的各个字符是否在 str2 里
func IsStr1InStr2(str1, str2 string) bool {
	m := map[rune]bool{}
	for _, v := range str2 {
		m[v] = true
	}
	for _, v := range str1 {
		if !m[v] {
			return false
		}
	}
	return true
}

// NonceStrCreate 生成随机字符串 相邻可重复
func NonceStrCreate2(len int, includeNumber bool) string {
	if len <= 0 {
		return ""
	}
	var bhint int
	var c rune
	var buffer bytes.Buffer
	var length int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		if includeNumber {
			if rdm := r.Intn(10); rdm < 3 {
				bhint = r.Intn(4) + r.Intn(23)
				c = rune(97 + bhint)
			} else if rdm < 6 {
				bhint = r.Intn(4) + r.Intn(23)
				c = rune(65 + bhint)
			} else {
				c = rune(48 + r.Intn(10))
			}
		} else {
			if r.Intn(10) <= 5 {
				bhint = r.Intn(4) + r.Intn(23)
				c = rune(97 + bhint)
			} else {
				bhint = r.Intn(4) + r.Intn(23)
				c = rune(65 + bhint)
			}
		}
		buffer.WriteString(string(c))
		length++
		if length == len {
			break
		}
	}
	return buffer.String()
}
