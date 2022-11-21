package base64

import (
	"strings"
)

const base64Chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(s string) string {
	res := []byte{}

	// if len(s) is not divisible by 3, then pad it with zero chars
	c := len(s) % 3
	padding := ""
	if c != 0 {
		for i := c; i < 3; i++ {
			s += string(rune(0))
			padding += "="
		}
	}

	for i := 0; i < len(s); i += 3 {
		n := int(s[i])<<(8*2) + int(s[i+1])<<(8*1) + int(s[i+2])
		res = append(res, base64Chars[(n>>(6*3))&63])
		res = append(res, base64Chars[(n>>(6*2))&63])
		res = append(res, base64Chars[(n>>(6*1))&63])
		res = append(res, base64Chars[n&63])
	}

	n := len(res)
	return string(res[:n-len(padding)]) + padding
}

func Decode(s string) string {
	if len(s) == 0 {
		return s
	}

	sbytes := []byte(s)
	padding := 0
	if s[len(s)-1] == '=' {
		sbytes[len(s)-1] = 'A'
		padding += 1
	}
	if s[len(s)-2] == '=' {
		sbytes[len(s)-2] = 'A'
		padding += 1
	}

	res := []byte{}
	for i := 0; i < len(sbytes); i += 4 {
		n := strings.IndexByte(base64Chars, sbytes[i])<<(6*3) +
			strings.IndexByte(base64Chars, sbytes[i+1])<<(6*2) +
			strings.IndexByte(base64Chars, sbytes[i+2])<<(6*1) +
			strings.IndexByte(base64Chars, sbytes[i+3])
		res = append(res, byte((n>>(8*2))&255))
		res = append(res, byte((n>>(8*1))&255))
		res = append(res, byte(n&255))
	}

	n := len(res)
	return string(res[:n-padding])
}
