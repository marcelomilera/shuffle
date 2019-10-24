package main

import (
	"fmt"
	"math/rand"
	"time"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	shuffled_array := shuffle(numbers)
	fmt.Println(shuffled_array)
}

/* 
Implementation of the "inside-out" version of the
Fisher-Yates shuffle algorithm 
Source: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_%22inside-out%22_algorithm
*/
func shuffle(numbers []int) []int {
	//Initializing the default Source to a deterministic state.
	rand.Seed(time.Now().UTC().UnixNano())
	
	for i := 0; i < len(numbers); i++ {
		j := rand.Intn(i+1)
		if j != i {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	return numbers
}