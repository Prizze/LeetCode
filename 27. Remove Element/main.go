package main

import "fmt"

func removeElement(nums []int, val int) int {
    for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]... )
			i--;
		}
	}
	return len(nums)
}

func main() {
	nums := []int{3,2,2,3}
	k := removeElement(nums, 3)

	fmt.Println(k)
	fmt.Println(nums)
}