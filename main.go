package main

import (
	"fmt"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	shuffled_array := shuffle(numbers)
	fmt.Println(shuffled_array)
}