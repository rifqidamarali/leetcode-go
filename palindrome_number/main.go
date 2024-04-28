package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// original := x
	reverse := 0
	for i := x; i >= 1; i = int(math.Floor(float64(i) / 10)) {
		reverse = reverse*10 + i%10
	}

	return reverse == x
}

func main() {
	fmt.Println(isPalindrome(1111)) 
}