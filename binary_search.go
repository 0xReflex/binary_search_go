package main

import (
	"fmt"
)

func main() {
	inputs_from_the_user()
}

func inputs_from_the_user() {
	var target, length_of_array, i int64
	var list_numbers []int64
	fmt.Println("what is the length of the array:")
	fmt.Println("Don't put large numbers or it will kill your pc and then you have to switch it off")
	fmt.Scanln(&length_of_array)
	fmt.Println("what is your target:")
	fmt.Scanln(&target)
	for i = 0; i <= length_of_array; i++ {
		list_numbers = append(list_numbers, i)
		}
	fmt.Println(binary_search(list_numbers,target))
}



// see how_it_works to know how this program works :) 
func binary_search(array []int64, target int64) (int, string, int){
	hops := 0
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right)/2
		if array[mid] == target {
			return mid , "Number of hops:", hops
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		hops += 1
	}
	return -1,"Number of hops:", hops
}
