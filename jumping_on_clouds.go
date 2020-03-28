package main

import "fmt"

func jumpingOnClouds(c []int) (jumps int) {
	currentPosition := 0
	for currentPosition < len(c)-1 {
		if currentPosition+2 <= len(c) && c[currentPosition+2] == 0 {
			fmt.Println("jumping", 2)
			currentPosition += 2
		} else {
			fmt.Println("jumping", 1)
			currentPosition += 1
		}
		jumps += 1
	}
	fmt.Println(jumps)
	return jumps
}

func main() {
	ar1 := []int{0, 1, 0, 0, 0, 1, 0}
	jumpingOnClouds(ar1)

	ar2 := []int{0, 0, 1, 0, 0, 1, 0}
	jumpingOnClouds(ar2)

	ar3 := []int{0, 0, 0, 0, 1, 0}
	jumpingOnClouds(ar3)
}
