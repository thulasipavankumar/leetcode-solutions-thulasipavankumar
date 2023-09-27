package main

import "fmt"

func reverseVowels(s string) string {
	//97, 101, 105, 111 117
	//65, 69, 73, 79,  85
	runes := []rune(s)
	low, high := 0, len(runes)-1
	for low < high {
		if !isVowel(runes[low]) {
			low++
			continue
		}
		if !isVowel(runes[high]) {
			high--
			continue
		}
		runes[low], runes[high] = runes[high], runes[low]
		low++
		high--

	}

	return string(runes)
}
func isVowel(r rune) bool {
	return (r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U')
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
