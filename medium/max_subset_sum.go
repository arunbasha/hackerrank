package main

import (
	"fmt"
	"math"
)

func maxSubsetSum(arr []int32) int32 {

	len := len(arr)
	if len == 0 {
		return math.MinInt32
	}

	if len == 1 {
		return arr[0]
	}

	arr[1] = max(arr[0], arr[2])
	for i := 2 ;i < len; i++ {
		tmp := max(arr[i-1], arr[i] + arr[i-2])
		arr[i] = max(arr[i], tmp)
	}

	return arr[len-1]
}

func max(x, y int32) int32{
	if x > y {
		return x
	}
	return y
}

func  main()  {
	fmt.Println(maxSubsetSum([]int32{-2,1,3,-4,5,}) ) // 8
	fmt.Println(maxSubsetSum([]int32{6,-2,-1,10,-3,-4,-5}) ) // 10
	fmt.Println(maxSubsetSum([]int32{-2,-1,-3,-4,5,9}) ) // 8
}