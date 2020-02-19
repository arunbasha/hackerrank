package main

import (
	"fmt"
	"strings"
)

func abbreviation(a string, b string) string {
	n := len(a)
	m := len(b)
	res  := make([][]int, m+1 )
	for i, _ := range res {
		res[i] = make([]int, n+1)
	}

	fmt.Println(res)

	for i:=0; i < n; i++ {
		for j := 0; j<= m; j++ {
			if res[i][j] != 1{
				continue
			}
			if j < m && strings.ToUpper(string(a[i])) == string(b[j]) {
				fmt.Printf("i=%d,j=%d\n", i+1, j+1)
				res[i+1][j+1] = 1
				fmt.Printf("res[%d][%d] = %d\n", i+1, j+1, res[i+1][j+1])
			}
			if strings.ToUpper(string(a[i])) == string(a[i]) {
				res [i+1][j] = 1
			}
		}
		//fmt.Println(res)
	}

	if res[n][m] == 1 {
		return "YES"
	}

	return "NO"
}

func  main()  {
	fmt.Println(abbreviation("AbcDE", "ABDE") )// "YES
	//fmt.Println(abbreviation("AbcDE", "AFDE")) // "NO
	//fmt.Println(abbreviation("daBcd", "ABC")) // "YES"
}