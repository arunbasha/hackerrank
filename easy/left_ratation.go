package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	length := int32(len(a))
	if d >= length {
		d = d % length
	}

	var res []int32
	for i := d; i<length; i++ {
		res = append(res, a[i])
	}
	for i := int32(0); i< d; i++ {
		res = append(res, a[i])
	}

	return res
}


func  main()  {
	fmt.Println(rotLeft([]int32{1,2,3,4,5}, 4) )// {5,1,2,3,4}
	fmt.Println(rotLeft([]int32{1,2,3,4,5}, 10) )// {5,1,2,3,4}
	fmt.Println(rotLeft([]int32{1,2,3,4,5}, 2) )// {3,4,5,1,2}
}