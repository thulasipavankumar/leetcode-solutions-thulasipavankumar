package main
import "fmt"

func main(){
	fmt.Println("Hello World!")
	fmt.Printf("%v,%v",isPalindrome(121),isPalindrome(567865))
}
func isPalindrome(x int) bool {
	if x < 0 
    str := fmt.Sprintf("%d",x)
	reverse := Reverse(str)
	return (str==reverse)
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}