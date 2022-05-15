// 判定是否互为字符重排
package main

import(

)


func main() {
	CheckPermutation('aabc', 'bcaa')
}

func CheckPermutation(s1 string, s2 string) bool {
	s1Len := strings.Count(s1,"")
	s2Len := strings.Count(s2,"")
	// 如果长度不同，直接返回FALSE
	if s1Len != s2Len {
		return false
	}
	// 定义一个map key-string value-int
	// var mapLit = map[string] int
 	mapList := make(map[string] int)
 	// 遍历字符串
 	for _, char := range s1 {
 		ch := string(char)
 		mapList[ch] ++
 	}
 	for _, char := range s2 {
 		ch := string(char)
 		mapList[ch] --
 	}
 	// 遍历map 如果值不全都为0，说明两个字符串不匹配
 	for _, value := range mapList {
 		if value != 0 {
 			return false
 		}
 	}
 	return true

}