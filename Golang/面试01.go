// 面试题 01.01. 判定字符是否唯一
package main

import(
	'fmt'
	'strings'
)
func main(){
	isUnique("leetcode")
}

func isUnique(astr string) bool {
	str := ""
	for _, char := range astr {
		if (strings.Contains(str, string(char))) {
			// 如果包含
			return false
		} else {
			str += string(char)
		}
	}
	return true
}