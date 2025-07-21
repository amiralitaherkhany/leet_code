package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(52))
}

func convertToTitle(columnNumber int) string {
	if columnNumber < 27 {
		return string(('A' + (columnNumber - 1)))
	}
	if columnNumber%26 == 0 {
		return convertToTitle((columnNumber/26)-1) + convertToTitle(26)
	}
	return convertToTitle(columnNumber/26) + convertToTitle(columnNumber%26)
}
