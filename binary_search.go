package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	target := 18
	fmt.Println(binary_search(a,target))
}
func binary_search(array []int, target int) int{
	left := 0
	right := len(array) - 1
	for left <= right {
		var mid int = (left + right)/2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
			// fmt.Println("we are at left:", left)
		} else {
			right = mid - 1
			// fmt.Println("we are at right:", right)
		}
	}
	return -1
}