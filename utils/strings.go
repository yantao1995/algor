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
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
