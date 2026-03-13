package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	k := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			k++
		}
	}

	// Swap elements in nums such that the first k elements are not equal to val
	for i := n - 1; i > k-1; i-- {
		if nums[i] == val {
			continue
		} else {
			for j := i - 1; j >= 0; j-- {
				if nums[j] == val {
					nums[j] = nums[i]
					nums[i] = val
					break
				}
			}
		}
	}

	return k
}

func modify(nums []int) {
	nums[0] = 100
}

func reassign(nums []int) {
	nums = []int{4, 5, 6}
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(nums, 2))
	fmt.Println(nums)
	println("--------------------")
	testArray := []int{1, 2, 3}
	modify(testArray)
	fmt.Println(testArray)
	reassign(testArray)
	fmt.Println(testArray)
}
