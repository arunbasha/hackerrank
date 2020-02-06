package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	cols := len(arr[0])
	rows := len(arr)
	prevSum := int32(-2147483648)
	sum:= int32(-2147483648)
	for i:= 0; i < rows - 2; i++ {
		for j:= 0; j < cols - 2; j++ {
			var curSum int32
			curSum += arr[i][j] + arr[i][j+1] +  arr[i][j+2]
			curSum += arr[i+1][j+1]
			curSum += arr[i+2][j] +arr[i+2][j+1] + arr[i+2][j+2]
			sum = curSum
			if prevSum > sum {
				sum = prevSum
			}
			prevSum = sum
		}
	}
	return sum
}


func main() {
	fmt.Println(hourglassSum([][]int32{
		{-9,-9,-9,1,1,1},
		{0,-9,0,4,3,2},
		{-9,-9,-9,1,2,3},
		{0,0,8,6,6,0},
		{0,0,0,-2,0,0},
		{0,0,1,2,4,0},
	}))
	fmt.Println(hourglassSum([][]int32{
		{1,1,1,0,0,0},
		{0,1,0,0,0,0},
		{1,1,1,0,0,0},
		{0,0,2,4,4,0},
		{0,0,0,2,0,0},
		{0,0,1,2,4,0},
	}))
	fmt.Println(hourglassSum([][]int32{
	{-1,-1,0,-9,-2,-2},
	{-2,-1,-6,-8,-2,-5},
	{-1,-1,-1,-2,-3,-4},
	{-1,-9,-2,-4,-4,-5},
	{-7,-3,-3,-2,-9,-9},
	{-1,-3,-1,-2,-4,-5},
	}))

}