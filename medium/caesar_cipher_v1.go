package main

import "fmt"
//Cipher main.go
func main(){
	var length, key int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &key)

	var res []rune
	for _, ch := range input {
		res = append(res, cipher(ch, key))
	}
	fmt.Println(string(res))
}

func cipher(r rune, key int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', key)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', key)
	}
	return r
}

func rotate(r rune, base, key int) rune {
	tmp := int(r) - base
	tmp = (tmp + key) % 26
	return rune(tmp + base)
}

