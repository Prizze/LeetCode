package main

import "fmt"



func findLucky(arr []int) int {
	intMap := map[int]int{}

	for _, n := range arr {
		intMap[n]++
	}

	var result int = -1
	for key, value := range intMap {
		if key == value && key > result {
			result = key
		}
	}
	return result
}

func main() {
	arr := []int{2,2,2,3,3}
	fmt.Println(findLucky(arr))
}