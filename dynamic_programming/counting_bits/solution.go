package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := range result {
		result[i] = len(strings.ReplaceAll(fmt.Sprintf("%b", i), "0", ""))
	}
	return result
}
