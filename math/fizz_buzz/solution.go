package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	slice := make([]string, n)

	for i := 0; i < n; i++ {

		if (i+1)%5 == 0 && (i+1)%3 == 0 {
			slice[i] += "FizzBuzz"
			continue
		}
		if (i+1)%5 == 0 {
			slice[i] += "Buzz"
			continue
		}
		if (i+1)%3 == 0 {
			slice[i] += "Fizz"
			continue
		}
		slice[i] = strconv.Itoa((i + 1))
	}

	return slice
}
