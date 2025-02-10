/**
 * @Author QG
 * @Date  2025/2/7 23:28
 * @description
**/

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode md5加密
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(password string, salt string) string {
	return MD5Encode(password + salt)
}

// ValidPassword 验证密码
func ValidPassword(plainpwd, salt string, password string) bool {
	md := Md5Encode(plainpwd + salt)
	return md == password
}
