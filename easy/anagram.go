package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abc", "bdc"))
	fmt.Println(isAnagram("abdc", "bac"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	firstArr := returnAlphaArr(s)
	secondArr := returnAlphaArr(t)
	//fmt.Println(firstArr,secondArr)
	for i := 0; i < len(firstArr); i++ {
		if firstArr[i] != secondArr[i] {
			return false
		}
	}

	return true
}
func returnAlphaArr(s string) [26]int {
	var result [26]int
	for i := 0; i < len(s); i++ {
		result[s[i]-97] += 1
	}
	return result

}
