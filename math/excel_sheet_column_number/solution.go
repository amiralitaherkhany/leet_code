package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(columnTitle string) int {
	if len(columnTitle) <= 1 {
		return int(columnTitle[0] - 64)
	}
	c := columnTitle[:len(columnTitle)-1]
	return titleToNumber(c)*26 + titleToNumber(string(columnTitle[len(columnTitle)-1]))
}
