package main

func hasDuplicate(nums []int) bool {
	var seen map[int]struct{} = make(map[int]struct{})
	// Look through entire list at most once. Time: O(n)
	for _, n := range nums {
		if _, exists := seen[n]; exists {
			return true
		}
		// Create at most n entries into the hashmap. Space: O(n)
		seen[n] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2}
	hasDuplicate(nums)
	println(hasDuplicate(nums))
}
