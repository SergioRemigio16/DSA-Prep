package main

func twoSum(nums []int, target int) []int {
	nums_map := make(map[int]int)
	// Add nums to a hashmap of (key int, count int)
	// Time: O(n)
	for _, n := range nums {
		nums_map[n]++
	}
	// for every number in nums
	// Time: O(n)
	for first_number_index, first_number := range nums {
		// second_number = target - firstNumber
		second_number := target - first_number
		// if hashmap contains secondNumber
		// Time: O(1)
		if _, ok := nums_map[second_number]; ok {
			second_number_index := 0
			// Look through nums array for secondNumber
			// Time: O(2N) Only happens at most twice
			for i, n := range nums {
				// Make sure same index doesn't return an answer
				if n == second_number {
					if first_number_index != i {
						second_number_index = i
						return []int{first_number_index, second_number_index}
					}
				}
			}
			// Return [firstNumber Index, secondNumber Index]
		} else {
			// Continue
		}
	}
	return []int{}
}

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	println(twoSum(nums, target)[0])
	println(twoSum(nums, target)[1])
	nums = []int{4, 5, 6}
	target = 10
	println(twoSum(nums, target)[0])
	println(twoSum(nums, target)[1])
	nums = []int{5, 5}
	target = 10
	println(twoSum(nums, target)[0])
	println(twoSum(nums, target)[1])
	nums = []int{1, 3, 4, 2}
	target = 6
	println(twoSum(nums, target)[0])
	println(twoSum(nums, target)[1])
}
