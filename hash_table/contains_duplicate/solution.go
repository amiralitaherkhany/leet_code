package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{2, 2}))
	fmt.Println(containsDuplicate([]int{2, 1, 3}))
}

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			return true
		}
		hashMap[v] = v
	}
	return false
}
