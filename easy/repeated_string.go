package main

import (
	"fmt"
)

func repeatedString(s string, n int64) int64 {
	var counter, resCounter int64
	len := len(s)
	for j:=0; j<len ; j++ {
		if 'a' == s[j] {
			counter++
		}
	}

	resCounter = (n/int64(len)) * counter

	rem := n % int64(len)

	for i:= int64(0); i<rem; i++ {
		if 'a' == s[i] {
			resCounter++
		}
	}
	return resCounter
}
func main() {
	fmt.Println(repeatedString("aba", 10)) // 7
	fmt.Println(repeatedString("a", 1000000000000)) // 1000000
	fmt.Println(repeatedString("abcac", 10)) // 4
	fmt.Println(repeatedString("abcac", 5)) // 4
	fmt.Println(repeatedString("abcac", 3)) // 4
}