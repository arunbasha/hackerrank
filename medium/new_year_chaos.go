package main

import "fmt"

// O(N)
func minimumBribes(q []int32) {
	var bribes, first, second, third int32
	first = 1
	second = 2
	third = 3
	len := len(q)

	for i:=0; i<len; i++ {
		if q[i] == first {
			first = second
			second = third
			third++
		} else if q[i] == second {
			bribes++
			second = third
			third++
		} else if q[i] == third {
			bribes += 2
			third++
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}

	fmt.Println(bribes)
}

// O(N2)
func minimumBribesV2(q []int32) {
	for i, v := range q {
		if v - int32(1) - int32(i) > int32(2) {
			fmt.Println("Too chaotic")
			return
		}
	}
	len := len(q)
	var swap int
	for i := 0; i<len-1; i++ {
		for j :=i+1; j<len; j++ {
			if q[i] > q[j] {
				tmp := q[j]
				q[j] = q[i]
				q[i] = tmp
				swap++
			}
		}
	}
	fmt.Println(swap)
}


func  main()  {
	minimumBribes([]int32{8,5,1,2,3,7,8,6,4})// "Too chaotic"
	minimumBribes([]int32{1,2,5,3,7,8,6,4}) // 7
	minimumBribes([]int32{2,1,5,3,4}) // 3
	minimumBribes([]int32{2,5,1,3,4}) // "Too chaotic"
}