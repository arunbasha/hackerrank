package main

import "fmt"

func minimumSwaps(arr []int32) int32 {
	len := int32(len(arr))
	var swaps int32
	for i := int32(0); i<len; {
		if arr[i] != i+1 {
			tmp := arr[i]
			arr[i] = arr[tmp-1]
			arr[tmp-1] = tmp
			swaps++
		}else {
			i++
		}
	}
	return swaps
}

func  main()  {
	fmt.Println(minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}) )// 5
	fmt.Println(minimumSwaps([]int32{2,3,4,1,5})) //3
	fmt.Println(minimumSwaps([]int32{4,3,1,2})) // 3
	fmt.Println(minimumSwaps([]int32{1,3,5,2,4,6,7})) // 3

}