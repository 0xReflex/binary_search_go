package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {
	inputs_from_the_user()
}

func inputs_from_the_user() { // It will take inputs from the user
	var target, length_of_array int
	fmt.Println("what is the length of the array:")
	fmt.Scanln(&length_of_array)
	fmt.Println("what is your target:")
	fmt.Scanln(&target)
	binary_search(remove_dublicate_from_array(the_sorted_array(random_array(length_of_array))), target)
}

func random_array(length_of_array int) []int { // this will make random array
	var i int
	var random_array []int
	for true{
		if i == length_of_array{
			return random_array
		}else {
			rand.Seed(time.Now().UnixNano() * int64(i))
			random_array = append(random_array ,rand.Intn(length_of_array))
			i++ 
		}
	}
	return random_array
}


func remove_dublicate_from_array(random_array []int) []int{ // It will remove dulicate values 
	safe_array ,map_of_the_array := []int{} ,make(map[int]bool)
	for _, entry := range random_array { 
		if _, value := map_of_the_array[entry]; !value { 
			map_of_the_array[entry] = true
			safe_array = append(safe_array , entry) 
		} 
	}
	return safe_array
}


func the_sorted_array(sorted_array []int) []int{//this will make sorted array
	sort.Ints(sorted_array)	
	return sorted_array
} 


// see how_it_works to know how this program works :) 
func binary_search(array []int, target int) int{
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
