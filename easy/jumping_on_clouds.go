package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	var count int32
	len := len(c)
	for i:=0; i< len - 1 ;{
		if i+2 < len && c[i+2] != 1 {
			count++
			i += 2
		} else {
			count++
			i++
		}
	}
	return count
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0,0,0,1,0,0})) //3
	fmt.Println(jumpingOnClouds([]int32{0,0,1,0,0,1,0})) //4
	fmt.Println(jumpingOnClouds([]int32{0,0,1,0,0,1,0,0,0,1,0})) //6
	fmt.Println(jumpingOnClouds([]int32{0,0,1,0,0,1,0,0,1,0})) //6
	fmt.Println(jumpingOnClouds([]int32{0,0,1,0,0,1,0,0,0})) //5
}