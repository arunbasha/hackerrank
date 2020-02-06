package main

import (
	"fmt"
	"strings"
)

func main(){
	var len, key int
	var input string
	fmt.Scanf("%d", &len)
	fmt.Scanf("%s", &input)
	fmt.Scanf("%d", &key)

	//Cipher main_2.go
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var res []rune

	for _, ch := range input {
		if strings.IndexRune(alphabetUpper, ch) > 0 {
			res = append(res, cipher(ch, []rune(alphabetUpper), key))
		} else if strings.IndexRune(alphabetLower, ch) >= 0 {
			res = append(res, cipher(ch, []rune(alphabetLower), key))
		} else {
			res = append(res, ch)
		}
	}
	fmt.Println(string(res))
}

func cipher(r rune, charset []rune, key int) rune{

	index := strings.IndexRune(string(charset), r)
	index = index + key
	if index > len(charset) {
		index = index % len(charset)
	}

	// Native approach
	//index := -1
	//for i, ch := range charset {
	//	if ch == r {
	//		index = i
	//	}
	//}
	//
	//if index < 0 {
	//	panic("Index < 0")
	//}
	//
	//for i:=0 ; i< key; i++ {
	//	index++
	//	if index > len(charset) {
	//		index = 0
	//	}
	//}

	return charset[index]
}