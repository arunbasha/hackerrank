package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	var level, valleyCounter int32

	for _, ch := range s {
		if ch == 'U' {
			level++
		}
		if ch == 'D' {
			level--
		}

		if level == 0 && ch == 'U' {
			valleyCounter++
		}
	}

	return valleyCounter
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU")) //1
	fmt.Println(countingValleys(8, "UDDDUDUUUDDU")) //2
	fmt.Println(countingValleys(8, "UDDDUDUUDDUUDU")) //3
	fmt.Println(countingValleys(8, "UDDDUDUUDUDUUUDDDU")) //4
}