package main

import "fmt"

func main() {
	fmt.Println(isAnagram("als", "lia"))
}

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	hashMap := make(map[rune]uint16)
	for _, c := range s {
		if _, ok := hashMap[c]; ok {
			hashMap[c] = hashMap[c] + 1
		} else {
			hashMap[c] = 1
		}

	}
	for _, c := range t {
		if _, ok := hashMap[c]; !ok {
			return false
		}
		hashMap[c] = hashMap[c] - 1
		if hashMap[c] == 0 {
			delete(hashMap, c)
		}
	}
	return true
}
