package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	inputs_from_the_user()
	// fmt.Println("total time taken: ", time.Since(t0))
}

func inputs_from_the_user() { // It will take inputs from the user
	var target, length_of_array int
	fmt.Println("what is the length of the array:")
	// fmt.Println("Don't put large numbers or it will kill your pc and then you have to switch it off")
	fmt.Scanln(&length_of_array)
	fmt.Println("what is your target:")
	fmt.Scanln(&target)
	binary_search(random_array(length_of_array), target)
}

func random_array(length_of_array int) []int { // this will make random array
	var i int
	var random_array []int
	for true{
		if i == length_of_array{
			return random_array
		}else {
			rand.Seed(time.Now().UnixNano())
			random_array = append(random_array ,rand.Intn(length_of_array))
			i++ 
		}
	}
	return random_array
}

// func the_sorted_array(sorted_array []int, target int){} //this will make sorted array
// see how_it_works to know how this program works :) 
func binary_search(array []int, target int) int{//(int, string, int)
	hops := 0
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right)/2
		if array[mid] == target {
			fmt.Println(mid , "Number of hops:", hops)
			return 0
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		hops += 1
	}
	fmt.Println( "Not found!!!!", "Number of hops:", hops)
	return -1 
}
