package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
    sockCount := make(map[int32]int32, n)

    for _, val := range ar {
        sockCount[val]++
    }
    var counter int32
    for _, val := range sockCount {
        if val/2 > 0 {
            counter += val / 2
        }
    }
    return counter
}

func main() {
    fmt.Println(sockMerchant(9, []int32{10,20,20,10,10,30,50,10,20}))
}

