package main

import (
	"strconv"
	"fmt"
)

func fizzbuzz(number int) string {
	if number % 15 == 0 { return "FizzBuzz" }
	if number % 5 == 0 { return "Buzz" }
	if number % 3 == 0 { return "Fizz" }
	return strconv.Itoa(number)
}

func main () {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
